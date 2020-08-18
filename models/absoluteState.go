package models



type AbsoluteState struct {

	Id string `json:"id"`
	Turn string `json:"turn"`

	State Player `json:"state"`


}

type Player struct{

	P1 PlayerState `json:"P1"`
	P2 PlayerState `json:"P2"`
	P3 PlayerState `json:"P3"`
	P4 PlayerState `json:"P4"`
	P5 PlayerState `json:"P5"`
	P6 PlayerState `json:"P6"`

}
type PlayerState struct{

UserName string `json:"username"`
Cards []string `json:"cards"`

}