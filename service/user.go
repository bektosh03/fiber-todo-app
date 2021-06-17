package service

//func (s *Service) CreateUser(req requests.CreateUser) (responses.User, error) {
//	req.Password = strings.TrimSpace(req.Password)
//	req.Email = strings.TrimSpace(req.Email)
//	ok, err := req.Validate()
//	if err != nil {
//		s.logger.Println("Error while validating:", err)
//		return responses.User{}, err
//	}
//	id, err := uuid.NewUUID()
//	if err != nil {
//		s.logger.Println("Error while generating UUID:", err)
//		return responses.User{}, err
//	}
//	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	req.FirstName = strings.TrimSpace(req.FirstName)
//	req.LastName = strings.TrimSpace(req.LastName)
//}
