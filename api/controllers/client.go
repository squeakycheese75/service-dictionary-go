package controllers

// func GetHomePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome to my protected home page!")
// }

// func  GetUser(id int) (*model.User, error) {
// req, err := c.newRequest("GET", fmt.Sprintf("/users/%d", id), nil)
// if err != nil {
// 	return nil, err
// }
// var user model.User
// _, err = c.do(req, &user)

// if err != nil {
// 	return nil, ErrUnavailable
// }

// return &user, err
// 	vars := mux.Vars(req)
// 	id := vars["id"]

// 	var source data.Source
// 	if result := env.DB.First(&source, id); result.Error != nil {
// 		utils.RespondWithError(res, http.StatusNotFound, result.Error.Error())
// 		return
// 	}
// 	utils.RespondWithJSON(res, http.StatusOK, source)

// }
// func WithCorrelationID(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		uuid := uuid.New()
// 		w.Header().Set("X-Api-Correlation-Id", uuid.String())
// 		h.ServeHTTP(w, r)
// 	}
// }

// func commonMiddleware(f http.HandlerFunc) http.HandlerFunc {
// 	return WithCorrelationID(f)
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	resBody, _ := json.Marshal(userRepository.GetUsers())
// 	w.Write(resBody)
// }
