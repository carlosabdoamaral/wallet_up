package utils

func CheckErr(err error, isFatal bool, msg string) error {
	if err != nil {
		if isFatal {

		}

		return err
	}

	return nil
}
