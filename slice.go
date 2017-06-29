package util

func DiffInt(new,old []int)(add ,del []int)  {
	f:=0 //already matched item's index in old slice
	add=make([]int,0)

	for _,n:=range new {
		var notMatch =true
		for i:=f;i<len(old) ;i++  {
			if n==old[i]{
				notMatch=false
				old[i],old[f]=old[f],old[i]
				f++
				break
			}
		}
		if notMatch{
			add=append(add,n)
		}
	}
	del=old[f:len(old)]
	return
}
