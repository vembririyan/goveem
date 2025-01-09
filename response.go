package goveem

func ResponseGet(res []map[string]interface{}, status int) map[string]interface{} {
	switch status {
	case 200:
		return map[string]interface{}{
			"status":  200,
			"message": "ok",
			"data":    res,
		}
	case 401:
		return map[string]interface{}{
			"status":  403,
			"message": "unauthorize",
			"data":    res,
		}
	case 403:
		return map[string]interface{}{
			"status":  403,
			"message": "forbidden",
			"data":    res,
		}
	case 404:
		return map[string]interface{}{
			"status":  404,
			"message": "not found",
		}
	default:
		return map[string]interface{}{
			"status":  500,
			"message": "internal Server error",
		}
	}
}

func ResponseModify(status int) map[string]interface{} {
	switch status {
	case 200:
		return map[string]interface{}{
			"status":  200,
			"message": "successfully",
		}
	case 201:
		return map[string]interface{}{
			"status":  201,
			"message": "created",
		}
	case 400:
		return map[string]interface{}{
			"status":  400,
			"message": "bad request",
		}
	case 405:
		return map[string]interface{}{
			"status":  405,
			"message": "method not allowed",
		}
	default:
		return map[string]interface{}{
			"status":  500,
			"message": "internal server error",
		}
	}
}
