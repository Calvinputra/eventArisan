package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GetRecid used to get inputter/authoriser recid identity from the jwt.
func GetRecid(c *gin.Context) (string, error) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
		return "", fmt.Errorf("Failed to get Authorization header")
	}
	auth := strings.Split(tokenString, " ")
	token, _, err := new(jwt.Parser).ParseUnverified(auth[1], jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Empty %s", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		inputter := fmt.Sprint(claims["recid"])
		return inputter, nil

	}
	return "", nil
}

func GetPartnerRecid(c *gin.Context) (string, error) {
	partnerRecid := c.Request.Header.Get("Partner-Recid")
	if partnerRecid == "" {
		return "", fmt.Errorf("Failed to get Partner-Recid")
	}
	return partnerRecid, nil
}

func GetBranchCodeFromHeader(ctx *gin.Context) (string, string) {
	branchRecid := ctx.Request.Header.Get("Branch-Code")

	isBranchCode := strings.Contains(branchRecid, ".")

	if isBranchCode {
		splitedBranchRecid := strings.Split(branchRecid, ".")
		if len(splitedBranchRecid) > 0 {
			return splitedBranchRecid[0], splitedBranchRecid[1]
		}
	}
	return branchRecid, ""
}

func GetBranchCodeFromJWT(ctx *gin.Context) []string {
	tokenString := ctx.Request.Header.Get("Authorization")

	re := regexp.MustCompile(`^Bearer\s.+$`)
	if tokenString == "" || !re.MatchString(tokenString) {
		fmt.Println("Empty authorization format is invalid")
		return nil
	}
	auth := strings.Split(tokenString, " ")
	// log.Printf("token string %s", auth[1])
	token, _, err := new(jwt.Parser).ParseUnverified(auth[1], jwt.MapClaims{})
	if err != nil {
		fmt.Println("Empty %s", err)
		return nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Empty: Couldn't extract claims")
		return nil
	}

	// TypeRecid assertion to get the value as a slice of interfaces
	branchCodeInterface, ok := claims["branchCodes"].([]interface{})
	if !ok {
		fmt.Println("Empty: Couldn't extract branchCodes as []interface{}")
		return nil
	}

	// Convert the slice of interfaces to a slice of strings
	var branchCode []string
	for _, recid := range branchCodeInterface {
		if recidStr, ok := recid.(string); ok {
			branchCode = append(branchCode, recidStr)
		} else {
			fmt.Println("Empty: Element in branchCode is not a string")
			return nil
		}
	}

	return branchCode
}
