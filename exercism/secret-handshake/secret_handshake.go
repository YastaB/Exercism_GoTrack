package secret

var responses =  [4]string{"wink", "double blink", "close your eyes", "jump"}
var indexes = [4]uint{1,2,4,8}
func Handshake(input uint) []string{
	var data []string
	if input & 16 != 0{
		for i:= len(responses) - 1; i >= 0 ; i--{
			if input & indexes[i] != 0{
				data = append(data, responses[i])
			}
		}
	}else{
		for i:= 0; i < len(responses) ; i++{
			if input & indexes[i] != 0{
				data = append(data, responses[i])
			}
		}
	}

	return data
}