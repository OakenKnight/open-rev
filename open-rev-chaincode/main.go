package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	// "github.com/hyperledger/fabric-chaincode-go/shim"
	// "os"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	domain "open-rev.com/domain"
)

type SmartContract struct {
	contractapi.Contract
}
type serverConfig struct {
	CCID    string
	Address string
}
const (
	timeFormat = "2006-01-02"
)

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	usersAssets := []domain.OpenRevUser{
		{
			ID: "helenab90453b68d0f40280391anisic", 
			Name: "Helena", 
			Surname: "Anisic", 
			Email: "hanisic@uns.ac.rs", 
			RoleId: 4, 
			Verified: true, 
			Code: "", 
			Type: "user",
			IsDeleted: false		},
		{ID: "14c3441fdcdf0cf67a7db7aaa9c81ffe", Name: "Aleksandar", Surname: "Ignjatijevic", Email: "alexignjat1998@gmail.com", RoleId: 3, Verified: true, Code: "", Type: "user", IsDeleted: false	},
		{ID: "123", Name: "PERA", Surname: "PERIC", Email: "blabla@gmail.com", RoleId: 3, Verified: true, Code: "", Type: "user", IsDeleted: false	},
		
	}

	roleAssets := []domain.Role{
		{ID: "3", Name: "Admin", Type: "role", IsDeleted: false},
		{ID: "4", Name: "Korisnik", Type: "role", IsDeleted: false},
	}

	areaAssets := []domain.Area{
		{ID: "25b03762e23dc89f15340520434b6cef", Name: "Elektronsko poslovanje", Hidden: false, Type: "area", IsDeleted: false},
		{ID: "4dbec17c53af916d7b2f7f1b0aa54255", Name: "Technology", Hidden: false, Type: "area", IsDeleted: false},
		{ID: "5c9de78f9e88ccacea97eac06793a68d", Name: "Education", Hidden: false, Type: "area", IsDeleted: false},
		{ID: "7af79e922519ac4be5c298f1814ae8d4", Name: "Science", Hidden: false, Type: "area", IsDeleted: false},
		{ID: "d7a2447e8b01299932fe54e632575e0f", Name: "Mathematics", Hidden: false, Type: "area", IsDeleted: false},
	}

	subareaAssets := []domain.SubArea{
		{ID: "255371cb6de81963bbfdcfdc6787dec8", Name: "Geometry", AreaId: "d7a2447e8b01299932fe54e632575e0f", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "2805e12cbdb7d8d1dd9b4a6a6165f2b9", Name: "Internet of Things", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "2962cf7ca98d9d3c975215b16b5ecae4", Name: "Elektronsko poslovanje", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "2bb26241046b5e6131e7c81443d5b360", Name: "Calculus", AreaId: "d7a2447e8b01299932fe54e632575e0f", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "2eb53a29978c5c3131e4f1e062dedb64", Name: "Vestacka inteligencija", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "31284ea7eb99871b405571dc929461c2", Name: "Robotics", AreaId: "4dbec17c53af916d7b2f7f1b0aa54255", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "4d378fba64c7aa1b1f2cdeeb4f2c0ab0", Name: "Big data", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "556ecbc01d3ed70d066a7d58ff65ee86", Name: "Physics", AreaId: "7af79e922519ac4be5c298f1814ae8d4", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "5639ee90bc391c8859bd5b48ac96834d", Name: "E-obrazovanje", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "5c7b2df95368cb53df5122881c8d81ab", Name: "Online", AreaId: "5c9de78f9e88ccacea97eac06793a68d", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "68decd3e2be58a8f421d7607de5edd80", Name: "Internet tehnologije", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "82901f7abe6b5f22e34ba68f7971ec6b", Name: "Chemistry", AreaId: "7af79e922519ac4be5c298f1814ae8d4", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "90bef33d8068382e34a65138cd67dd77", Name: "Internet marketing i drustveni mediji", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "a1754617725241311638f7e1725d2edf", Name: "Algebra", AreaId: "d7a2447e8b01299932fe54e632575e0f", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "a6dae2d72c256fd128295370623fe2e3", Name: "Softversko inzenjerstvo", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "c00553a7ac6abeb701d32ae2f99630e6", Name: "Information Technology", AreaId: "4dbec17c53af916d7b2f7f1b0aa54255", Hidden: false, Type: "subarea", IsDeleted: false},
		{ID: "d4c7018a1cfe7dadb054e80fa5b1f73b", Name: "Cloud computing", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: true},
		{ID: "f6fab332d25a2af4b306a1ac4e5b80c8", Name: "Mobilno poslovanje", AreaId: "25b03762e23dc89f15340520434b6cef", Hidden: false, Type: "subarea", IsDeleted: true},
	}

	sciworkAssets := []domain.ScientificWork{
		{ID: "0bf60b871e8ff560573deb7c4c7d673f", Title: "Prvi rad o vestackoj inteligenciji", SubAreaId: "2eb53a29978c5c3131e4f1e062dedb64", Abstract: "prvi testni abstract", Keywords: "kljucne reci o ovom radu", PdfFile: "testopenrev/a1271408-c68e-11ed-9139-bce92f86d338_t1.csv", UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", Type: "scientific-work", PublishDate: "2022-06-22", IsDeleted: false},
		{ID: "22d2e594f8239d49ae5851d09583756a", Title: "Drugi rad", SubAreaId: "2eb53a29978c5c3131e4f1e062dedb64", Abstract: "drugi testni abstract", Keywords: "kljucne reci o ovom radu", PdfFile: "testopenrev/dd6e56b0-ca0e-11ed-9629-bce92f86d338_t2.csv", UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", Type: "scientific-work", PublishDate: "2021-09-11", IsDeleted: false},
	}

	reviewAssets := []domain.Review{
		{ID: "067b502089fbf82146e4bf6b879326f2", Review: "Veoma dobra prica", Assessment: 5, Recommend: true, UserId: "helenab90453b68d0f40280391anisic", ScientificWorkId: "0bf60b871e8ff560573deb7c4c7d673f", Type: "review", IsDeleted: false},
		{ID: "0f3fbf4ab0062a2089a1b0cd37a855fb", Review: "Losa prica", Assessment: 1, Recommend: false, UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", ScientificWorkId: "0bf60b871e8ff560573deb7c4c7d673f", Type: "review", IsDeleted: false},

		{ID: "13aeb21ca61481b3ab56154c7319470c	", Review: "Veoma dobra prica broj 3", Assessment: 4, Recommend: true, UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", ScientificWorkId: "22d2e594f8239d49ae5851d09583756a", Type: "review", IsDeleted: false},
	}

	reviewQualityAssets := []domain.ReviewQuality{
		{ID: "f980356e3549c0c7b4f7b804bae70a5f", Assessment: 5, UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", ReviewId: "067b502089fbf82146e4bf6b879326f2", Type: "review-quality", IsDeleted: false},

		{ID: "aca9ab5f8eade613af0685df3af6f26b", Assessment: 1, UserId: "helenab90453b68d0f40280391anisic", ReviewId: "0f3fbf4ab0062a2089a1b0cd37a855fb", Type: "review-quality", IsDeleted: false},
		{ID: "aca9ab5f8eade613af0685df3af6f23b", Assessment: 3, UserId: "14c3441fdcdf0cf67a7db7aaa9c81ffe", ReviewId: "0f3fbf4ab0062a2089a1b0cd37a855fb", Type: "review-quality", IsDeleted: false},

		{ID: "d199e0e139ddac1696b4b35d8a1acd79	", Assessment: 3, UserId: "helenab90453b68d0f40280391anisic", ReviewId: "13aeb21ca61481b3ab56154c7319470c", Type: "review-quality", IsDeleted: false},
	}

	for _, role := range roleAssets {
		roleJSON, err := json.Marshal(role)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(role.ID, roleJSON)
		if err != nil {
			return fmt.Errorf("failed to put open-rev-role to world state %v", err)
		}
	}

	for _, user := range usersAssets {
		userJSON, err := json.Marshal(user)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(user.ID, userJSON)
		if err != nil {
			return fmt.Errorf("failed to put open-rev-user to world state. %v", err)
		}

		//indexName := "email~roleid~id"
		//emailRoleID, err := ctx.GetStub().CreateCompositeKey(indexName, []string{user.Email, strconv.Itoa(user.RoleId), user.ID})
		//if err != nil {
		//	return err
		//}
		//
		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(emailRoleID, value)
		//if err != nil {
		//	return err
		//}
	}

	for _, area := range areaAssets {
		areaJSON, err := json.Marshal(area)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(area.ID, areaJSON)
		if err != nil {
			return fmt.Errorf("failed to put area to world state. %v", err)
		}

		//indexName := "area_name~id"
		//areaNameId, err := ctx.GetStub().CreateCompositeKey(indexName, []string{area.Name, area.ID})
		//if err != nil {
		//	return err
		//}
		//
		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(areaNameId, value)
		//if err != nil {
		//	return err
		//}

	}
	for _, subarea := range subareaAssets {
		subareaJSON, err := json.Marshal(subarea)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(subarea.ID, subareaJSON)
		if err != nil {
			return fmt.Errorf("failed to put subarea to world state. %v", err)
		}

		//indexName := "subarea_name~area_id~id"
		//subareaNameAreaIdId, err := ctx.GetStub().CreateCompositeKey(indexName, []string{subarea.Name, subarea.AreaId, subarea.ID})
		//if err != nil {
		//	return err
		//}
		//
		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(subareaNameAreaIdId, value)
		//if err != nil {
		//	return err
		//}

	}

	for _, sciwork := range sciworkAssets {
		sciworkJSON, err := json.Marshal(sciwork)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(sciwork.ID, sciworkJSON)
		if err != nil {
			return fmt.Errorf("failed to put scientific-work to world state. %v", err)
		}

		//indexName := "title~publishdate~userId~subareaId~ID"
		//titleDateUserIdSubareaIdID, err := ctx.GetStub().CreateCompositeKey(indexName, []string{sciwork.Title, sciwork.PublishDate, sciwork.UserId, sciwork.SubAreaId, sciwork.ID})
		//if err != nil {
		//	return err
		//}
		//
		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(titleDateUserIdSubareaIdID, value)
		//if err != nil {
		//	return err
		//}
	}

	for _, review := range reviewAssets {
		reviewJSON, err := json.Marshal(review)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(review.ID, reviewJSON)
		if err != nil {
			return fmt.Errorf("failed to put review to world state. %v", err)
		}

		//indexName := "userId~scientificWorkId"
		//userSciWork, err := ctx.GetStub().CreateCompositeKey(indexName, []string{review.UserId, review.ScientificWorkId})
		//if err != nil {
		//	return err
		//}

		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(userSciWork, value)
		//if err != nil {
		//	return err
		//}

	}

	for _, reviewQuality := range reviewQualityAssets {
		reviewQualityJSON, err := json.Marshal(reviewQuality)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(reviewQuality.ID, reviewQualityJSON)
		if err != nil {
			return fmt.Errorf("failed to put review to world state. %v", err)
		}

		//indexName := "userId~reviewId"
		//userReview, err := ctx.GetStub().CreateCompositeKey(indexName, []string{reviewQuality.UserId, reviewQuality.ReviewId})
		//if err != nil {
		//	return err
		//}
		//
		//value := []byte{0x00}
		//err = ctx.GetStub().PutState(userReview, value)
		//if err != nil {
		//	return err
		//}

	}

	return nil
}
func (s *SmartContract) CreateRevUserAsset(ctx contractapi.TransactionContextInterface, id, name, surname, code, email string) (*domain.OpenRevUser, error) {
	users, err := s.ReadUserByEmail(ctx, email)
	if len(users) != 0 {
		return nil, fmt.Errorf("User with email %s already exists", email)
	}
	dto := domain.OpenRevUser{ID: id, Name: name, Surname: surname, Email: email, RoleId: 4, Verified: false, Code: code, Type: "user", IsDeleted : false}

	userJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	log.Println(userJSON)

	err = ctx.GetStub().PutState(dto.ID, userJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put open-rev user to world state. %v", err)

	}
	//
	//indexName := "email~roleid~id"
	//emailRoleID, err := ctx.GetStub().CreateCompositeKey(indexName, []string{dto.Email, strconv.Itoa(dto.RoleId), dto.ID})
	//if err != nil {
	//	return nil, err
	//}
	//
	//value := []byte{0x00}
	//err = ctx.GetStub().PutState(emailRoleID, value)
	//if err != nil {
	//	return nil, err
	//}
	return &dto, nil
}

func (s *SmartContract) CreateReviewAsset(ctx contractapi.TransactionContextInterface, id, sciId, userId, assessment, recommend, review string) (*domain.Review, error) {
	sciWork, err := s.ReadScientificWorkAsset(ctx, sciId)
	if err != nil {
		return nil, err
	}
	if sciWork.UserId == userId {
		return nil, fmt.Errorf("user %s cant review his own paper with id %s", userId, sciId)
	}
	reviewers, err := s.GetAllReviewersOnScientificWork(ctx, sciId)
	if err != nil {
		return nil, err
	}
	for _, r := range reviewers {
		if r.ID == userId {
			return nil, fmt.Errorf("user %s already reviewed scientific work %s", userId, sciId)
		}
	}

	a, err := strconv.Atoi(assessment)
	if err != nil {
		return nil, err

	}
	rec, err := strconv.ParseBool(recommend)

	dto := domain.Review{ID: id, ScientificWorkId: sciId, Assessment: a, Recommend: rec, Review: review, 
		UserId: userId, Type: "review", IsDeleted: false}

	reviewJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	log.Println(reviewJSON)

	err = ctx.GetStub().PutState(dto.ID, reviewJSON)
	if err != nil {
		return nil,	fmt.Errorf("failed to put review to world state. %v", err)

	}
	//
	//indexName := "userId~scientificWorkId"
	//userSciWork, err := ctx.GetStub().CreateCompositeKey(indexName, []string{dto.UserId, dto.ScientificWorkId})
	//if err != nil {
	//	return nil, err
	//}
	//
	//value := []byte{0x00}
	//err = ctx.GetStub().PutState(userSciWork, value)
	//if err != nil {
	//	return nil, err
	//}

	return &dto, nil
}


func (s *SmartContract) CreateAreaAsset(ctx contractapi.TransactionContextInterface, id, name string) (*domain.Area, error) {
	
	dto := domain.Area{ID: id, Name: name, Hidden: false, Type: "area", IsDeleted: false}

	areaJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	log.Println(areaJSON)

	err = ctx.GetStub().PutState(dto.ID, areaJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put area to world state. %v", err)

	}
	
	return &dto, nil
}

func (s *SmartContract) CreateSubAreaAsset(ctx contractapi.TransactionContextInterface, id, areaId, name string) (*domain.SubArea, error) {
	
	dto := domain.SubArea{ID: id, Name: name, AreaId: areaId ,Hidden: false, Type: "subarea", IsDeleted: false}

	_, err := s.AreaAssetExists(ctx,areaId)
	if err!=nil {
		return nil, fmt.Errorf("area asset does not exist");
	}


	subareaJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	log.Println(subareaJson)

	err = ctx.GetStub().PutState(dto.ID, subareaJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put subarea to world state. %v", err)

	}
	
	return &dto, nil
}






func (s *SmartContract) CreateReviewQualityAsset(ctx contractapi.TransactionContextInterface, id, reviewId, userId, assessment string) (*domain.ReviewQuality, error) {
	rev, err := s.ReadReviewAsset(ctx, reviewId)
	if err != nil {
		return nil, err
	}
	if rev.UserId == userId {
		return nil, fmt.Errorf("user %s cant rate his own review with id %s", userId, reviewId)
	}
	reviewers, err := s.GetAllReviewersOnReview(ctx, reviewId)
	if err != nil {
		return nil, err
	}
	for _, r := range reviewers {
		if r.ID == userId {
			return nil, fmt.Errorf("user %s has already reviewed review with id %s", userId, reviewId)
		}
	}
	a, err := strconv.Atoi(assessment)
	if err != nil {
		return nil, err

	}

	dto := domain.ReviewQuality{ID: id, ReviewId: reviewId, Assessment: a, UserId: userId, Type: "review-quality", IsDeleted : false}

	reviewJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}
	log.Println(reviewJSON)

	err = ctx.GetStub().PutState(dto.ID, reviewJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put review to world state. %v", err)

	}



	
	return &dto, nil
}

func (s *SmartContract) VerifyRevUserAsset(ctx contractapi.TransactionContextInterface, code, id string) (*domain.OpenRevUser, error) {
	exists, err := s.OpenRevUserAssetExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	openRevUserAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}
	if openRevUserAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", id)
	}

	var openRevUserAsset domain.OpenRevUser
	err = json.Unmarshal(openRevUserAssetJSON, &openRevUserAsset)
	if err != nil {
		return nil, err
	}

	if openRevUserAsset.Verified {
		return nil, fmt.Errorf("The open-rev user has already been verified")
	}
	if openRevUserAsset.Code != code {
		return nil, fmt.Errorf("wrong code supplied for verification")
	}

	dto := domain.OpenRevUser{ID: id, Name: openRevUserAsset.Name, Surname: openRevUserAsset.Surname, Email: openRevUserAsset.Email, RoleId: openRevUserAsset.RoleId, Verified: true, Code: "", Type: "user"}

	userJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, userJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put open-rev user to world state. %v", err)

	}

	return &openRevUserAsset, nil
}
func (s *SmartContract) GetAllReviewersOnScientificWork(ctx contractapi.TransactionContextInterface, sciWorkId string) ([]*domain.OpenRevUser, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.OpenRevUser
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Review
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "review" && sciWorkId == asset.ScientificWorkId {

			userAsset, err := s.ReadOpenRevUserAsset(ctx, asset.UserId)
			if err != nil {
				return nil, err
			}

			assets = append(assets, userAsset)
		}
	}

	return assets, nil
}
func (s *SmartContract) GetAllReviewersOnReview(ctx contractapi.TransactionContextInterface, reviewId string) ([]*domain.OpenRevUser, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.OpenRevUser
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "review-quality" && reviewId == asset.ReviewId {

			userAsset, err := s.ReadOpenRevUserAsset(ctx, asset.UserId)
			if err != nil {
				return nil, err
			}

			assets = append(assets, userAsset)
		}
	}

	return assets, nil
}

//
//func (s *SmartContract) IsReviewRated(ctx contractapi.TransactionContextInterface, userId, reviewId string) (bool, error) {
//	worksIter, err := ctx.GetStub().GetStateByPartialCompositeKey("userId~reviewId", []string{userId, reviewId})
//	if err != nil {
//		return true, err
//	}
//
//	defer worksIter.Close()
//
//	if worksIter.HasNext() {
//		return true, nil
//	}
//	return false, nil
//}

//
//func (s *SmartContract) IsSciWorkReviewed(ctx contractapi.TransactionContextInterface, userId, sciWorkId string) (bool, error) {
//	worksIter, err := ctx.GetStub().GetStateByPartialCompositeKey("userId~scientificWorkId", []string{sciWorkId})
//	if err != nil {
//		return true, err
//	}
//
//	defer worksIter.Close()
//
//	k := 0
//	for i := 0; worksIter.HasNext(); i++ {
//		responseRange, err := worksIter.Next()
//		if err != nil {
//			return true, err
//		}
//
//		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
//		if err != nil {
//			return true, err
//		}
//
//		userid := compositeKeyParts[0]
//		if userid == userId {
//			k++
//		}
//
//	}
//
//	return k != 0, nil
//}

func (s *SmartContract) ReadScientificWorkAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.ScientificWork, error) {
	sciWorkAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read scientific work from world state: %v", err)
	}
	if sciWorkAssetJSON == nil {
		return nil, fmt.Errorf("the scientific work asset %s does not exist", id)
	}

	var sciWorkAsset domain.ScientificWork
	err = json.Unmarshal(sciWorkAssetJSON, &sciWorkAsset)
	if err != nil {
		return nil, err
	}

	return &sciWorkAsset, nil
}

func (s *SmartContract) CreateScientificWorkAsset(ctx contractapi.TransactionContextInterface, id, title, abstract, keywords, pdf, subareaId, userId, date string) (*domain.ScientificWork, error) {
	//worksWithSameName, err := s.ReadAllScientificWorksByTitle(ctx, title)

	//if err != nil {
	//	return nil, err
	//}

	//if len(worksWithSameName) != 0 {
	//	return nil, fmt.Errorf("Scientific paper with same name already exists!")
	//}
	subareaExists, err := s.SubAreaAssetExists(ctx, subareaId)
	if err != nil {
		return nil, err
	}
	if !subareaExists {
		return nil, fmt.Errorf("the subarea asset %s does not exits", subareaId)
	}
	userExists, err := s.OpenRevUserAssetExists(ctx, userId)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, fmt.Errorf("the open-rev-user asset %s does not exist", id)
	}

	newWork := domain.ScientificWork{ID: id, SubAreaId: subareaId, Title: title, Abstract: abstract, Keywords: keywords, 
		PdfFile: pdf, UserId: userId, PublishDate: date, Type: "scientific-work", IsDeleted : false}

	newWorkJSON, err := json.Marshal(newWork)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(newWork.ID, newWorkJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put scientific-work user to world state. %v", err)

	}
	//
	//indexName := "title~publishdate~userId~subareaId~ID"
	//titleDateUserIdSubareaIdID, err := ctx.GetStub().CreateCompositeKey(indexName, []string{newWork.Title, newWork.PublishDate, newWork.UserId, newWork.SubAreaId, newWork.ID})
	//if err != nil {
	//	return nil, err
	//}
	//
	//value := []byte{0x00}
	//err = ctx.GetStub().PutState(titleDateUserIdSubareaIdID, value)
	//if err != nil {
	//	return nil, err
	//}

	return &newWork, nil
}

func (s *SmartContract) EditScientificWorkAsset(ctx contractapi.TransactionContextInterface, id, title, abstract, keywords, pdf, subareaId, userId, date string) (*domain.ScientificWork, error) {

	exists, err := s.ScientificWorkAssetExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	subareaExists, err := s.SubAreaAssetExists(ctx, subareaId)
	if err != nil {
		return nil, err
	}
	if !subareaExists {
		return nil, fmt.Errorf("the subarea asset %s does not exits", subareaId)
	}

	userExists, err := s.OpenRevUserAssetExists(ctx, userId)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, fmt.Errorf("the open-rev-user asset %s does not exist", id)
	}

	//worksWithSameName, err := s.ReadAllScientificWorksByTitle(ctx, title)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(worksWithSameName) != 0 {
	//	return nil, fmt.Errorf("Scientific paper with same name already exists!")
	//}
	//// maybe a check for the same file, but that has to be discussed later

	scientificWorkAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read scientific work asset from world state: %v", err)
	}
	if scientificWorkAssetJSON == nil {
		return nil, fmt.Errorf("the scientific work asset %s does not exist", id)
	}

	var scientificWorkAsset domain.ScientificWork
	err = json.Unmarshal(scientificWorkAssetJSON, &scientificWorkAsset)
	if err != nil {
		return nil, err
	}
	sciWork := domain.ScientificWork{ID: id, Title: title, PublishDate: date, Abstract: abstract, Keywords: keywords, 
		PdfFile: pdf, UserId: userId, SubAreaId: subareaId, Type: "scientific-work", IsDeleted: false}

	sciWorkJSON, err := json.Marshal(sciWork)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(sciWork.ID, sciWorkJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put scientific-work to world state. %v", err)

	}
	//indexName := "title~publishdate~userId~subareaId~ID"
	//titleDateUserIdSubareaIdID, err := ctx.GetStub().CreateCompositeKey(indexName, []string{sciWork.Title, sciWork.PublishDate, sciWork.UserId, sciWork.SubAreaId, sciWork.ID})
	//if err != nil {
	//	return nil, err
	//}
	//
	//value := []byte{0x00}
	//err = ctx.GetStub().PutState(titleDateUserIdSubareaIdID, value)
	//if err != nil {
	//	return nil, err
	//}

	return &sciWork, nil
}

func (s *SmartContract) DeleteAreaAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.Area, error){
	exists, err := s.AreaAssetExists(ctx, id)
	if err!=nil {
		return nil, err
	}
	if !exists{
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	areaAssetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read area from world state: %v", err)
	}

	if areaAssetJson == nil {
		return nil, fmt.Errorf("the area asset %s does not exist", id)
	}

	var area domain.Area
	err = json.Unmarshal(areaAssetJson, &area)
	if err != nil {
		return nil, err
	}

	dto := domain.Area{ID: area.ID, Name: area.Name, Hidden: area.Hidden, Type: area.Type, IsDeleted : true}

	areaJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, areaJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put area to world state. %v", err)

	}

	return &area, nil
}





func (s *SmartContract) DeleteSubAreaAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.SubArea, error){
	exists, err := s.SubAreaAssetExists(ctx, id)
	if err!=nil {
		return nil, err
	}
	if !exists{
		return nil, fmt.Errorf("the asset &s does not exist", id)
	}

	subAreaAssetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read subarea from world state: %v", err)
	}

	if subAreaAssetJson == nil {
		return nil, fmt.Errorf("the subarea asset %s does not exist", id)
	}

	var subarea domain.SubArea
	err = json.Unmarshal(subAreaAssetJson, &subarea)
	if err != nil {
		return nil, err
	}

	dto := domain.SubArea{ID: subarea.ID, Name: subarea.Name, Hidden: subarea.Hidden, Type: subarea.Type, IsDeleted : true, AreaId: subarea.AreaId}

	subareaJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, subareaJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put subarea to world state. %v", err)

	}

	return &subarea, nil
}




func (s *SmartContract) DeleteOpenRevUserAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.OpenRevUser, error){
	exists, err := s.OpenRevUserAssetExists(ctx, id)
	if err!=nil {
		return nil, err
	}
	if !exists{
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	openrevuserAssetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read openrevUser from world state: %v", err)
	}

	if openrevuserAssetJson == nil {
		return nil, fmt.Errorf("the openrevUser asset %s does not exist", id)
	}

	var user domain.OpenRevUser
	err = json.Unmarshal(openrevuserAssetJson, &user)
	if err != nil {
		return nil, err
	}

	dto := domain.OpenRevUser{	ID:user.ID, Name: user.Name, Surname: user.Name, Email: user.Email, RoleId: user.RoleId, 
		Verified: user.Verified, Code: user.Code, Type: user.Type, IsDeleted: true}

	userJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, userJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put openrevUser to world state. %v", err)

	}

	return &user, nil
}



func (s *SmartContract) EditOpenRevUserAsset(ctx contractapi.TransactionContextInterface, id, name, surname string) (*domain.OpenRevUser, error) {
	exists, err := s.OpenRevUserAssetExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	openRevUserAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}

	if openRevUserAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", id)
	}

	var openRevUserAsset domain.OpenRevUser
	err = json.Unmarshal(openRevUserAssetJSON, &openRevUserAsset)
	if err != nil {
		return nil, err
	}

	dto := domain.OpenRevUser{ID: id, Name: name, Surname: surname, Email: openRevUserAsset.Email, RoleId: openRevUserAsset.RoleId, 
		Verified: openRevUserAsset.Verified, Type: "user", IsDeleted: openRevUserAsset.IsDeleted}

	userJSON, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, userJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to put open-rev user to world state. %v", err)
	}

	return &openRevUserAsset, nil
}


func (s *SmartContract) OpenRevUserAssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	openRevUserAsset, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return openRevUserAsset != nil, nil
}

//
//func (s *SmartContract) IsAlreadyReviewedScientificWork(ctx contractapi.TransactionContextInterface, userId, sciWorkId string) (bool, error) {
//	res, err := ctx.GetStub().GetState(userId + "~" + sciWorkId)
//	if err != nil || res != nil {
//		return true, err
//	}
//	return false, nil
//}
//
//func (s *SmartContract) IsAlreadyReviewedReview(ctx contractapi.TransactionContextInterface, userId, reviewId string) (bool, error) {
//	res, err := ctx.GetStub().GetState(userId + "~" + reviewId)
//	if err != nil || res != nil {
//		return true, err
//	}
//
//	return false, nil
//}

func (s *SmartContract) ReviewExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	review, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return review != nil, nil
}

func (s *SmartContract) ReviewQualityExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	reviewQ, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return reviewQ != nil, nil
}

func (s *SmartContract) SubAreaAssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	subarea, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return subarea != nil, nil
}

func (s *SmartContract) AreaAssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	area, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return area != nil, nil
}

func (s *SmartContract) ScientificWorkAssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	sciWorkAsset, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return sciWorkAsset != nil, nil
}

func (s *SmartContract) ReadOpenRevUserAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.OpenRevUser, error) {
	openRevUserAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}
	if openRevUserAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", id)
	}

	var openRevUserAsset domain.OpenRevUser
	err = json.Unmarshal(openRevUserAssetJSON, &openRevUserAsset)
	if err != nil {
		return nil, err
	}

	return &openRevUserAsset, nil
}

func (s *SmartContract) ReadReviewQaualityAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.ReviewQuality, error) {
	reviewQualityJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read review-quality from world state: %v", err)
	}
	if reviewQualityJSON == nil {
		return nil, fmt.Errorf("the review-quality asset %s does not exist", id)
	}

	var reviewQualityAsset domain.ReviewQuality
	err = json.Unmarshal(reviewQualityJSON, &reviewQualityAsset)
	if err != nil {
		return nil, err
	}

	return &reviewQualityAsset, nil
}

func (s *SmartContract) ReadReviewAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.Review, error) {
	reviewJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read review from world state: %v", err)
	}
	if reviewJSON == nil {
		return nil, fmt.Errorf("the review asset %s does not exist", id)
	}

	var reviewAsset domain.Review
	err = json.Unmarshal(reviewJSON, &reviewAsset)
	if err != nil {
		return nil, err
	}

	return &reviewAsset, nil
}

func (s *SmartContract) ReadSubAreaAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.SubArea, error) {
	subAreaJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read subarea from world state: %v", err)
	}
	if subAreaJSON == nil {
		return nil, fmt.Errorf("the subarea asset %s does not exist", id)
	}

	var subAreaAsset domain.SubArea
	err = json.Unmarshal(subAreaJSON, &subAreaAsset)
	if err != nil {
		return nil, err
	}

	return &subAreaAsset, nil
}
func (s *SmartContract) ReadAreaAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.Area, error) {
	areaJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read area from world state: %v", err)
	}
	if areaJSON == nil {
		return nil, fmt.Errorf("the area asset %s does not exist", id)
	}

	var areaAsset domain.Area
	err = json.Unmarshal(areaJSON, &areaAsset)
	if err != nil {
		return nil, err
	}

	return &areaAsset, nil
}

func (s *SmartContract) ReadAllScientificWorkAssetsWithDetails(ctx contractapi.TransactionContextInterface) ([]*domain.ScientificWorkForSortingDTO, error){
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.ScientificWorkForSortingDTO
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ScientificWork
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "scientific-work" {
			var sciWorkDetails domain.ScientificWorkForSortingDTO
			err = json.Unmarshal(queryResponse.Value, &sciWorkDetails)
			if err != nil {
				return nil, err
			}
			sciWorkDetails.AvgRate, err = s.GetScientificWorkAvgMark(ctx, asset.ID)
			if err != nil {
				return nil, fmt.Errorf("failed to calculate average mark for scientific work %s", asset.ID)
			}
			user, err := s.ReadOpenRevUserAsset(ctx, asset.UserId)
			if err != nil {
				return nil, fmt.Errorf("failed to read user %s from world state ", asset.UserId)
			}
			sciWorkDetails.User = user.Name + " " + user.Surname
			assets = append(assets, &sciWorkDetails)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadOpenRevUserInfoAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.OpenRevUserInfoDTO, error) {
	openRevUserAssetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read person from world state: %v", err)
	}
	if openRevUserAssetJSON == nil {
		return nil, fmt.Errorf("the person asset %s does not exist", id)
	}

	var openRevUserAsset domain.OpenRevUserInfoDTO
	err = json.Unmarshal(openRevUserAssetJSON, &openRevUserAsset)
	if err != nil {
		return nil, err
	}

	sciWorksByUser, err := s.ReadAllScientificWorksByUserAssets(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Error getting sci works for user %s", id)
	}
	var totalMark float32
	totalMark = 0.0
	counter := 0
	for _, sciWork := range sciWorksByUser {
		avgMarkForSciWork, err := s.GetScientificWorkAvgMark(ctx, sciWork.ID)
		if err != nil {
			return nil, fmt.Errorf("Error getting sci work avg mark for sci work %s", sciWork.ID)
		}
		if avgMarkForSciWork != 0.0 {
			counter++
		}
		totalMark += avgMarkForSciWork
	}
	
	if totalMark == 0.0 || counter ==0 {
		openRevUserAsset.AvgMark = 0.0
	} else {
		openRevUserAsset.AvgMark = totalMark / float32(counter)
	}

	revsByUser, err := s.ReadAllReviewsByOpenRevUserAssets(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Error getting reviews for user %s", id)
	}

	totalRevMark := 0
	counter = 0

	for _, rev := range revsByUser {
		totalRevMark += rev.Assessment
		counter++
	}

	if totalRevMark == 0.0 || counter ==0 {
		openRevUserAsset.AvgReview = 0.0
	} else {
		openRevUserAsset.AvgReview = float32(totalRevMark) / float32(counter)
	}

	var totalRevQualityMark float32
	totalRevQualityMark = 0.0

	counter = 0
	allRevsQ, err := s.ReadAllReviewQualityAssets(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error getting all rev quality assets")
	}

	for _, rev := range revsByUser {
		for _, revQ := range allRevsQ {
			if revQ.ReviewId == rev.ID {
				totalRevQualityMark += float32(revQ.Assessment)
				counter++
			}
		}
	}

	if totalRevQualityMark == 0.0 || counter== 0 {
		openRevUserAsset.AvgRevQuality = 0.0
	} else {
		openRevUserAsset.AvgRevQuality = totalRevQualityMark / float32(counter)
	}

	userAvgMark, err := s.GetAverageQualityMarkForUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Error getting average quality for user %s", id)
	}

	openRevUserAsset.AvgMyRevsQuality = userAvgMark
	openRevUserAsset.ReviewsCount =  len(revsByUser)
	openRevUserAsset.WorksCount = len(sciWorksByUser)
	openRevUserAsset.IsAdmin = openRevUserAsset.RoleId == 3
	openRevUserAsset.ID = id
	return &openRevUserAsset, nil
}

func (s *SmartContract) ReadAllOpenRevUserAssets(ctx contractapi.TransactionContextInterface) ([]*domain.OpenRevUser, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.OpenRevUser
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.OpenRevUser
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "user" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllReviewAssets(ctx contractapi.TransactionContextInterface) ([]*domain.Review, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.Review
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Review
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "review" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllScientificWorkAssets(ctx contractapi.TransactionContextInterface) ([]*domain.ScientificWork, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.ScientificWork
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ScientificWork
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "scientific-work" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllScientificWorksByUserAssets(ctx contractapi.TransactionContextInterface, userId string) ([]*domain.ScientificWork, error) {

	exists, err := s.OpenRevUserAssetExists(ctx, userId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", userId)
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	//var assets []*domain.ScientificWork
	assets := make([]*domain.ScientificWork, 0)
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ScientificWork
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "scientific-work") && (asset.UserId == userId) {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadScientificWorkDetails(ctx contractapi.TransactionContextInterface, id string) (*domain.ScientificWorkDetailsDTO, error) {
	sciWorkJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read scientific work from world state: %v", err)
	}
	if sciWorkJSON == nil {
		return nil, fmt.Errorf("the sci work asset %s does not exist", id)
	}

	var sciWork domain.ScientificWork
	var reviewsDTOArray = make([]domain.ReviewForDetailsDTO, 0)

	var sciWorkDetails domain.ScientificWorkDetailsDTO
	err = json.Unmarshal(sciWorkJSON, &sciWork)
	if err != nil {
		return nil, err
	}
	//
	reviews, err := s.ReadAllReviewsByScientificPaperAssets(ctx, sciWork.ID)
	if err != nil {
		return nil, err
	}
	if reviews == nil {
		sciWorkDetails.Review = make([]domain.ReviewForDetailsDTO, 0)
	} else {
		// TODO: possible refactoring having in mind too many calls towards the ledger
		//
		for _, rev := range reviews {
			strAss := strconv.Itoa(rev.Assessment)
			strRec := strconv.FormatBool(rev.Recommend)
			user, err := s.ReadOpenRevUserAsset(ctx, rev.UserId)
			if err != nil {
				return nil, err
			}
			sumRevQ, err := s.GetSumOfRevQualityByReview(ctx, rev.ID)
			if err != nil {
				return nil, err
			}
			countRevQ, err := s.GetCountOfRevQualityByReview(ctx, rev.ID)
			if err != nil {
				return nil, err
			}
			reviewers, err := s.GetUsersOfRevQualityByReview(ctx, rev.ID)
			if err != nil {
				return nil, err
			}

			reviewDTO := domain.ReviewForDetailsDTO{ReviewId: rev.ID, Review: rev.Review, Assessment: strAss,
				Recommend: strRec, UserId: rev.UserId, User: user.Name + " " + user.Surname, ScientificWorkId: rev.ScientificWorkId,
				SumRevQuality: sumRevQ, CountRevQuality: countRevQ}

			if reviewers == nil {
				reviewDTO.UsersRevQuality = make([]string, 0)
			} else {
				reviewDTO.UsersRevQuality = reviewers

			}

			reviewsDTOArray = append(reviewsDTOArray, reviewDTO)
		}

		sciWorkDetails.Review = reviewsDTOArray

	}

	if len(reviews) == 0 {
		sciWorkDetails.AvgMark = 0.0
	} else {
		avg, err := s.GetScientificWorkAvgMark(ctx, id)
		if err != nil {
			return nil, err
		} else if avg == -1.0 {
			return nil, fmt.Errorf("invalid review avg mark")
		} else {
			sciWorkDetails.AvgMark = avg
		}
	}
	subarea, err := s.ReadSubAreaAsset(ctx, sciWork.SubAreaId)
	if err != nil {
		return nil, err
	}
	area, err := s.ReadAreaAsset(ctx, subarea.AreaId)
	if err != nil {
		return nil, err
	}
	sciWorkDetails.Area = area.Name + " / " + subarea.Name

	sciWorkDetails.WorkInfo = sciWork
	return &sciWorkDetails, nil
}

func (s *SmartContract) ReadUserByEmail(ctx contractapi.TransactionContextInterface, email string) ([]*domain.OpenRevUser, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	//var assets []*domain.OpenRevUser

	assets := make([]*domain.OpenRevUser, 0)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.OpenRevUser
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "user") && (asset.Email == email) {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllScientificWorksBySubAreaAssets(ctx contractapi.TransactionContextInterface, subareaId string) ([]*domain.ScientificWork, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.ScientificWork
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ScientificWork
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "scientific-work") && (asset.SubAreaId == subareaId) {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllReviewsByOpenRevUserAssets(ctx contractapi.TransactionContextInterface, id string) ([]*domain.Review, error) {
	exists, err := s.ScientificWorkAssetExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var assets []*domain.Review
	assets := make([]*domain.Review, 0)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Review
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "review") && (asset.UserId == id) {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllReviewsByScientificPaperAssets(ctx contractapi.TransactionContextInterface, id string) ([]*domain.Review, error) {
	exists, err := s.ScientificWorkAssetExists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// var assets []*domain.Review
	assets := make([]*domain.Review, 0)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Review
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "review" && asset.ScientificWorkId == id {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) GetSumOfRevQualityByReview(ctx contractapi.TransactionContextInterface, id string) (int, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return -1, err
	}
	defer resultsIterator.Close()
	sum := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return -1, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return -1, err
		}

		if (asset.Type == "review-quality") && (asset.ReviewId == id) {
			sum += asset.Assessment
		}

	}

	return sum, nil
}

func (s *SmartContract) GetCountOfRevQualityByReview(ctx contractapi.TransactionContextInterface, id string) (int, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return -1, err
	}
	defer resultsIterator.Close()
	num := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return -1, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return -1, err
		}

		if (asset.Type == "review-quality") && (asset.ReviewId == id) {
			num++
		}

	}

	return num, nil
}

func (s *SmartContract) GetUsersOfRevQualityByReview(ctx contractapi.TransactionContextInterface, id string) ([]string, error) {
	var usersArray []string
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "review-quality") && (asset.ReviewId == id) {
			usersArray = append(usersArray, asset.UserId)
		}
	}

	return usersArray, nil
}

func (s *SmartContract) GetScientificWorkAvgMark(ctx contractapi.TransactionContextInterface, id string) (float32, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return -1.0, err
	}
	defer resultsIterator.Close()
	num := 0
	sum := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return -1.0, err
		}

		var asset domain.Review
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return -1, err
		}
		if (asset.Type == "review") && (asset.ScientificWorkId == id) {
			sum += asset.Assessment
			num++
		}
	}
	if num == 0{
		return float32(0), nil
	}
	
	return float32((sum * 1.0) / (num * 1.0)), nil
}

func (s *SmartContract) ReadAllDashboardItemAssets(ctx contractapi.TransactionContextInterface) ([]*domain.DashboardItem, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.DashboardItem
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ScientificWork
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "scientific-work" {

			openRevUserAsset, err := s.ReadOpenRevUserAsset(ctx, asset.UserId)
			if err != nil {
				return nil, err
			}
			
			dashboardItem := domain.DashboardItem{ID: asset.ID, Title: asset.Title, Abstract: asset.Abstract, Keywords: asset.Keywords, 
				PdfFile: asset.PdfFile, User: openRevUserAsset.Name + " " + openRevUserAsset.Surname, PublishDate: asset.PublishDate}

			dashboardItem.AverageRate, err = s.GetScientificWorkAvgMark(ctx, asset.ID)
			if err != nil {
				return nil, fmt.Errorf("Failed to read average mark for sci work %s", asset.ID)
			}
			assets = append(assets, &dashboardItem)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllAreaSubareaAssets(ctx contractapi.TransactionContextInterface) ([]*domain.AreaSubareaDTO, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.AreaSubareaDTO
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Area
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "area" {
			subareas, err := s.ReadAllSubAreaByAreaIdAssets(ctx, asset.ID)
			if err != nil {
				return nil, err
			}
			dto := domain.AreaSubareaDTO{Area: asset.Name, SubAreas: subareas}
			assets = append(assets, &dto)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllAreaAssets(ctx contractapi.TransactionContextInterface) ([]*domain.Area, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.Area
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.Area
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "area" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllSubAreaByAreaIdAssets(ctx contractapi.TransactionContextInterface, areaId string) ([]domain.SubAreaDTO, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []domain.SubAreaDTO
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.SubArea
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if (asset.Type == "subarea") && (asset.AreaId == areaId) {
			dto := domain.SubAreaDTO{ID: asset.ID, SubArea: asset.Name, IsDeleted: asset.IsDeleted}
			assets = append(assets, dto)
		}
	}

	return assets, nil
}
func (s *SmartContract) ReadAllSubAreaAssets(ctx contractapi.TransactionContextInterface) ([]*domain.SubArea, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.SubArea
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.SubArea
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "subarea" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}

func (s *SmartContract) ReadAllReviewQualityAssets(ctx contractapi.TransactionContextInterface) ([]*domain.ReviewQuality, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	assets := make([]*domain.ReviewQuality, 0)

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "review-quality" {
			assets = append(assets, &asset)
		}
	}

	return assets, nil
}



func (s *SmartContract) ReadAllUsersWithDetails(ctx contractapi.TransactionContextInterface) ([]*domain.OpenRevUserInfoDTO, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*domain.OpenRevUserInfoDTO
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset domain.OpenRevUser
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		if asset.Type == "user" {
			var openRevUserAsset domain.OpenRevUserInfoDTO
			err = json.Unmarshal(queryResponse.Value, &openRevUserAsset)
			if err != nil {
				return nil, err
			}
			openRevUserAsset.ID = asset.ID
			sciWorksByUser, err := s.ReadAllScientificWorksByUserAssets(ctx, asset.ID)
			if err != nil {
				return nil, fmt.Errorf("Error getting sci works for user %s", asset.ID)
			}
			var totalMark float32
			totalMark = 0.0
			counter := 0
		
			for _, sciWork := range sciWorksByUser {
				avgMarkForSciWork, err := s.GetScientificWorkAvgMark(ctx, sciWork.ID)
				if err != nil {
					return nil, fmt.Errorf("Error getting sci work avg mark for sci work %s", sciWork.ID)
				}
				if avgMarkForSciWork != 0.0 {
					counter++
				}
				totalMark += avgMarkForSciWork
			}
			if totalMark == 0.0 {
				openRevUserAsset.AvgMark = 0.0
			} else {
				openRevUserAsset.AvgMark = totalMark / float32(counter)
			}
		
			revsByUser, err := s.ReadAllReviewsByOpenRevUserAssets(ctx, asset.ID)
			if err != nil {
				return nil, fmt.Errorf("Error getting reviews for user %s", asset.ID)
			}
		
			totalRevMark := 0
			counter = 0
		
			for _, rev := range revsByUser {
				totalRevMark += rev.Assessment
				counter++
			}
			if totalRevMark == 0.0 || counter ==0{
				openRevUserAsset.AvgReview = 0.0
			} else {
				openRevUserAsset.AvgReview = float32(totalRevMark) / float32(counter)
			}
		
			var totalRevQualityMark float32
			totalRevQualityMark = 0.0
		
			counter = 0
			allRevsQ, err := s.ReadAllReviewQualityAssets(ctx)
			if err != nil {
				return nil, fmt.Errorf("Error getting all rev quality assets")
			}
			for _, rev := range revsByUser {
				for _, revQ := range allRevsQ {
					if revQ.ReviewId == rev.ID {
						totalRevQualityMark += float32(revQ.Assessment)
						counter++
					}
				}
			}
		
			if totalRevQualityMark == 0.0 || counter == 0{
				openRevUserAsset.AvgRevQuality = 0.0
			} else {
				openRevUserAsset.AvgRevQuality = totalRevQualityMark / float32(counter)
			}
		
			userAvgMark, err := s.GetAverageQualityMarkForUser(ctx, asset.ID)
			if err != nil {
				return nil, fmt.Errorf("Error getting average quality for user %s", asset.ID)
			}
		
			openRevUserAsset.AvgMyRevsQuality = userAvgMark
			openRevUserAsset.ReviewsCount = len(revsByUser)
			openRevUserAsset.WorksCount = len(sciWorksByUser)
			openRevUserAsset.IsAdmin = openRevUserAsset.RoleId == 3


			assets = append(assets, &openRevUserAsset)


		}

	}

	return assets, nil
}




func (s *SmartContract) GetAverageMarkForReview(ctx contractapi.TransactionContextInterface, revId string) (float32, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return -1.0, err
	}
	defer resultsIterator.Close()
	num := 0
	sum := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return -1.0, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return -1, err
		}
		if (asset.Type == "review-quality") && (asset.ReviewId == revId) {
			sum += asset.Assessment
			num++
		}
	}

	return float32((sum * 1.0) / (num * 1.0)), nil
}

func (s *SmartContract) GetAverageQualityMarkForUser(ctx contractapi.TransactionContextInterface, userId string) (float32, error) {

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return -1.0, err
	}
	defer resultsIterator.Close()
	num := 0
	sum := 0
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return -1.0, err
		}

		var asset domain.ReviewQuality
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return -1, err
		}
		if (asset.Type == "review-quality") && (asset.UserId == userId) {
			sum += asset.Assessment
			num++
		}
	}
	var retVal float32
	if num != 0 {
		retVal = float32((sum * 1.0) / (num * 1.0))
	}
	return retVal, nil
}

func (s *SmartContract) DeleteScientificWorkAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.ScientificWork, error){
	exists, err := s.ScientificWorkAssetExists(ctx, id)
	if err!=nil {
		return nil, err
	}
	if !exists{
		return nil, fmt.Errorf("the asset &s does not exist", id)
	}

	scientificWorkAssetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read scientificWork from world state: %v", err)
	}

	if scientificWorkAssetJson == nil {
		return nil, fmt.Errorf("the scientificWork asset %s does not exist", id)
	}

	var sci domain.ScientificWork
	err = json.Unmarshal(scientificWorkAssetJson, &sci)
	if err != nil {
		return nil, err
	}

	dto := domain.ScientificWork{ID: sci.ID, SubAreaId: sci.SubAreaId, Title: sci.Title, PublishDate: sci.PublishDate, Abstract: sci.Abstract, 
		Keywords: sci.Keywords, PdfFile: sci.PdfFile, UserId: sci.UserId, Type: sci.Type, IsDeleted : true}

	areaJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(dto.ID, areaJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put scientificWork to world state. %v", err)
	}

	return &sci, nil
}


func (s *SmartContract) DeleteReviewAsset(ctx contractapi.TransactionContextInterface, id string) (*domain.Review, error){
	reviewAssetJson, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read review from world state: %v", err)
	}

	if reviewAssetJson == nil {
		return nil, fmt.Errorf("the review asset %s does not exist", id)
	}
	
	var rev domain.Review
	err = json.Unmarshal(reviewAssetJson, &rev)
	if err != nil {
		return nil, err
	}


	dto := domain.Review{ID: id, Review: rev.Review, Assessment: rev.Assessment, Recommend: rev.Recommend, UserId: rev.UserId,  
		ScientificWorkId: rev.ScientificWorkId, Type: rev.Type, IsDeleted : true}

	reviewJson, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	err = ctx.GetStub().PutState(id, reviewJson)
	if err != nil {
		return nil, fmt.Errorf("failed to put review to world state. %v", err)
	}

	return &rev, nil
}

func main() {


	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creating open-rev chaincode: %v", err)
	}
	
	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting open-rev chaincode: %v", err)
	}



		// config := serverConfig{
		// 	CCID:    os.Getenv("CHAINCODE_ID"),
		// 	Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
		// }

		// chaincode, err := contractapi.NewChaincode(&SmartContract{})

		// if err != nil {
		// 	log.Panicf("error create openrev chaincode: %s", err)
		// }

		// server := &shim.ChaincodeServer{
		// 	CCID:    config.CCID,
		// 	Address: config.Address,
		// 	CC:      chaincode,
		// 	TLSProps: shim.TLSProperties{
		// 		Disabled: true,
		// 	},
		// }
		// if err := server.Start(); err != nil {
		// 	log.Panicf("error starting open-rev chaincode: %s", err)
		// }
	

}
