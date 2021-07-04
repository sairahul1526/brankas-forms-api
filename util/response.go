package util

import (
	"encoding/json"
	CONSTANT "forms-api/constant"
	"net/http"
	"strconv"
)

// SetReponse - set request response with status, message
func SetReponse(w http.ResponseWriter, status, errorMsg, msg, msgType string, response map[string]interface{}) {
	w.Header().Set("Status", status)
	response["meta"] = setMeta(status, errorMsg, msg, msgType)
	s, _ := strconv.Atoi(status)
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(response)
}

func setMeta(status, errorMsg, msg, msgType string) map[string]string {
	if len(msg) == 0 {
		if status == CONSTANT.StatusCodeBadRequest {
			msg = "Bad Request"
		} else if status == CONSTANT.StatusCodeServerError {
			msg = "Server Error"
		}
	}
	return map[string]string{
		"status":        status,
		"error_message": errorMsg,
		"message":       msg,
		"message_type":  msgType,
	}
}

func getHTTPStatusCode(code string) int {
	switch code {
	case CONSTANT.StatusCodeOk:
		return http.StatusOK
	case CONSTANT.StatusCodeCreated:
		return http.StatusCreated
	case CONSTANT.StatusCodeBadRequest:
		return http.StatusBadRequest
	case CONSTANT.StatusCodeServerError:
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
