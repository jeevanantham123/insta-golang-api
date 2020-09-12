package repo

import (
	"fmt"
	"log"

	"github.com/jeevanantham123/insta-golang-api/model"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//Signup func for new User
func Signup(db *gorm.DB, user model.User) (string, string) {

	var res model.User
	data := db.Where("Email = ?", user.Email).First(&res).RecordNotFound()

	if data == true {
		if err := db.Create(&user).Error; err != nil {
			return "", err.Error()
		}
		return "User added successfully", ""
	}

	return "", "Email already exists - Error"
}

//Login func
func Login(db *gorm.DB, username string, password string) (*gorm.DB, error) {

	var res model.User
	result := db.Where("user_name = ? AND password = ?", username, password).Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

//Friends func
func Friends(db *gorm.DB, username string) ([]string, error) {
	var res model.UserFriend
	var Friendsarray []string
	result, err := db.Select("friends").Where("user_name = ?", username).Find(&res).Rows()
	if err == nil {
		for result.Next() {
			er := result.Scan(pq.Array(&Friendsarray))
			if er != nil {
				return nil, er
			}
		}
		return Friendsarray, nil
	}
	return nil, err
}

//About func
func About(db *gorm.DB, username string) (string, error) {
	var res model.User
	var about string
	result, err := db.Select("about").Where("user_name = ?", username).Find(&res).Rows()
	if err == nil {
		for result.Next() {
			er := result.Scan(&about)
			if er != nil {
				return "", er
			}
		}
		return about, nil
	}
	return "", err
}

//Profile func
func Profile(db *gorm.DB, username string) (string, error) {
	var res model.User
	var ProfileURL string
	result, err := db.Select("profile_url").Where("user_name = ?", username).Find(&res).Rows()
	if err == nil {
		for result.Next() {
			er := result.Scan(&ProfileURL)
			if er != nil {
				return "", er
			}
		}
		return ProfileURL, nil
	}
	return "", err
}

//SuggestionTable func
func SuggestionTable(db *gorm.DB, username string) ([]model.SuggestionTab, error) {
	var suggestions []model.SuggestionTab = []model.SuggestionTab{}
	query := "SELECT friends FROM user_friends where user_name = ?"
	query2, dberr := db.Raw(query, username).Rows()
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	var userfriendsarray []string
	var friendoffriendsarray []string
	var suggestedarray []string
	row, erro := db.Raw("select user_name from friend_suggestions").Rows()
	if erro != nil {
		return nil, erro
	}
	var usr string
	for row.Next() {
		err := row.Scan(&usr)

		if err != nil {
			panic(err)
		}
		suggestedarray = append(suggestedarray, usr)
	}
	for query2.Next() {
		err := query2.Scan(pq.Array(&userfriendsarray))
		if err != nil {
			panic(err)
		}
		for i := range userfriendsarray {
			q, e := db.Raw("SELECT friends FROM user_friends where user_name = ?", userfriendsarray[i]).Rows()
			if e != nil {
				log.Fatalf("Unable to execute the query. %v", e)
			}
			for q.Next() {
				er := q.Scan(pq.Array(&friendoffriendsarray))
				if er != nil {
					return nil, er
				}
				fmt.Println(friendoffriendsarray)
				for j := range friendoffriendsarray {
					var flag = 0
					if friendoffriendsarray[j] != username {
						for k := range suggestedarray {
							if suggestedarray[k] == friendoffriendsarray[j] {
								flag = 1
							}
						}
						if flag == 0 {
							rr := db.Exec("Insert into friend_suggestions (user_name,followed) values (?,?)", friendoffriendsarray[j], false)
							if rr.Error != nil {
								return nil, rr.Error
							}
						}
					}
				}
			}
		}
	}
	rows, errr := db.Raw("select id,user_name,followed from friend_suggestions").Rows()
	if errr != nil {
		return nil, errr
	}
	for rows.Next() {
		suggestion := model.SuggestionTab{}
		err := rows.Scan(&suggestion.ID, &suggestion.UserName, &suggestion.Followed)
		row, dberr := db.Raw("SELECT profile_url from users where user_name = ?", suggestion.UserName).Rows()
		for row.Next() {
			er := row.Scan(&suggestion.ProfileURL)
			if dberr != nil || er != nil {
				return nil, er
			}
		}
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}
	err := rows.Err()
	if err != nil {
		return nil, err
	}
	fmt.Println(suggestions)
	return suggestions, nil
}

//Requesting func
func Requesting(db *gorm.DB, username string, friendname string) ([]string, error) {
	query := "Select requested from user_friends where user_name = ?"
	rows, dberr := db.Raw(query, friendname).Rows()
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	var requestedarray []string
	for rows.Next() {
		er := rows.Scan(pq.Array(&requestedarray))
		if er != nil {
			return nil, er
		}
	}
	requestedarray = append(requestedarray, username)
	er := db.Exec("Update user_friends set requested = ? where user_name = ?", pq.Array(requestedarray), friendname)
	if er.Error != nil {
		return nil, er.Error
	}
	return requestedarray, nil
}

//Accepting func
func Accepting(db *gorm.DB, username string, friendname string) ([]string, error) {
	type request struct {
		RequestedArray []string `json:"requestedarray"`
	}
	var req = request{}
	var AcceptedArray []string
	dbqu, dber := db.Raw("Select friends from user_friends where user_name = ?", username).Rows()
	if dber != nil {
		return nil, dber
	}
	for dbqu.Next() {
		er := dbqu.Scan(pq.Array(&AcceptedArray))
		if er != nil {
			panic(er)
		}
	}
	var FriendArray []string
	dbquF, dberF := db.Raw("Select friends from user_friends where user_name = ?", friendname).Rows()
	if dberF != nil {
		return nil, dberF
	}
	for dbquF.Next() {
		er := dbquF.Scan(pq.Array(&FriendArray))
		if er != nil {
			panic(er)
		}
	}
	FriendArray = append(FriendArray, username)
	query := "Select requested from user_friends where user_name = ?"
	rows, dberr := db.Raw(query, username).Rows()
	if dberr != nil {
		log.Fatalf("Unable to execute the query. %v", dberr)
	}
	for rows.Next() {
		er := rows.Scan(pq.Array(&req.RequestedArray))
		if er != nil {
			panic(er)
		}
	}
	for i := range req.RequestedArray {
		if req.RequestedArray[i] == friendname {
			AcceptedArray = append(AcceptedArray, friendname)
			req.RequestedArray[i] = req.RequestedArray[len(req.RequestedArray)-1]
			req.RequestedArray[len(req.RequestedArray)-1] = ""
		}
	}
	req.RequestedArray = req.RequestedArray[:len(req.RequestedArray)-1]
	er := db.Exec("Update user_friends set requested = ? where user_name = ?", pq.Array(req.RequestedArray), username)
	if er.Error != nil {
		return nil, er.Error
	}
	e := db.Exec("Update user_friends set friends = ? where user_name = ?", pq.Array(AcceptedArray), username)
	if e.Error != nil {
		return nil, e.Error
	}
	el := db.Exec("Update user_friends set friends = ? where user_name = ?", pq.Array(FriendArray), friendname)
	if el.Error != nil {
		return nil, el.Error
	}
	return req.RequestedArray, nil
}
