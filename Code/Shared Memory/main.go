package main

import (
	"bytes"
	"fmt"

	"github.com/hidez8891/shm"
)

func main() {

	w, _ := shm.Create("shm_name", 8196)
	defer w.Close()

	r, _ := shm.Open("shm_name", 8196)
	defer r.Close()

	wbuf := []byte(`6600.6006|362=0.46,3=1388|4=1388|5=IBM|8=157.73|9=100|1021=0|854=p|394=157.73|474=157.749977558|
	1584=89029|444=20160713|459=93001362639104|461=26142|55=93001362849000|22=89120|361=0.71|981=0|16=1468416601.372|
	18=1468416601.372|362=0.45,3=1388|4=1388|5=BAC|8=157.73|9=100|1021=0|854=p|394=157.73|474=157.749977558|1584=89029|
	444=20160713|459=93001362639104|461=26142|55=93001362849000|22=89120|361=0.71|981=0|16=1468416601.372
	|18=1468416601.372|362=0.45,3=1388|4=1388|5=IBM|8=157.69|9=100|1021=600|854=p|394=157.69|474=157.749910415|
	1584=89129|444=20160713|459=93004178882560|461=27052|55=93004179085000|22=89328|361=0.67|981=1|16=1468416604.1916|
	18=1468416604.1916|362=0.43,3=1388|4=1388|5=BAC|8=157.69|9=100|1021=600|854=p|394=157.69|474=157.749910415|1584=89129|
	444=20160713|459=93004178882560|461=27052|55=93004179085000|22=89328|361=0.67|981=1|16=1468416604.1916|
	18=1468416604.1916|362=0.43,3=1388|4=1388|5=IBM|8=157.72|9=100|1021=600|854=p|474=157.749876969|1584=89229|
	444=20160713|459=93004178888960|461=27117|55=93004179085000|22=89428|361=0.7|981=1|16=1468416604.1916|
	18=1468416604.1916|362=0.45,3=1388|4=1388|5=BAC|8=157.72|9=100|1021=600|854=p|474=157.749876969|1584=89229|
	444=20160713|459=93004178888960|461=27117|55=93004179085000|22=89428|361=0.7|981=1|16=1468416604.1916|18=1468416604.1916
	|362=0.45,3=1388|4=1388|5=IBM|8=157.7|9=100|1021=600|854=n|474=157.749843597|1584=89329|444=20160713|459=93004178916000|461=21|
	55=93004179085000|22=89528|361=0.7|981=0|16=1468416604.1916|
	18=1468416604.1916|362=0.45,`)
	w.Write(wbuf)

	rbuf := make([]byte, len(wbuf))
	r.Read(rbuf)

	rtemp := bytes.Split(rbuf, []byte(","))

	tokens := make(map[string]string)
	for v := range rtemp {

		resultbuf := rtemp[v]
		n := bytes.Index(resultbuf, []byte("|5="))
		m := bytes.Index(resultbuf, []byte("|8="))
		k := bytes.Index(resultbuf, []byte("|9="))

		if n != -1 && m != -1 {
			a := n + 3
			b := a + (m - n - 3)

			c := m + 3
			d := c + (k - m - 3)

			key := string(resultbuf[a:b])
			value := string(resultbuf[c:d])

			tokens[key] = value

			/*if _, ok := tokens[key]; ok {
				//delete(tokens, key)
				tokens[key] = value
			} else {
				tokens[key] = value
			}*/
		}
	}

	var IBMValue, BACVaule string
	fmt.Print("\033[s")
	for k, v := range tokens {
		fmt.Print("\033[u\033[K")
		if k == "IBM" {
			IBMValue = v
			fmt.Print("IBM : ", IBMValue)
			fmt.Print("\tBAC : ", BACVaule)
		} else if k == "BAC" {
			BACVaule = v
			fmt.Print("IBM : ", IBMValue)
			fmt.Print("\tBAC : ", BACVaule)
		} else {
			fmt.Print("IBM : ")
			fmt.Print("\tBAC : ")
		}
	}
}