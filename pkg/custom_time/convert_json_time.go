package customtime

import "time"

type CustomTime time.Time
const layout = "2006-01-02T15:04"

//converte json em time.Time
func(ct *CustomTime) UnmarshalJson(b []byte) error {
	s := string(b)
	t,err := time.Parse(layout, s[1:len(s)-1])
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}