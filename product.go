package main


func parseProductPage(seed string) {
	
	/*
	_, err := url.ParseRequestURI(args[0])

	if err != nil {
		fmt.Println(err.Error())
	}
	*/



} // parseProductPage


func productJob() {

	tb := Tb{}

	//tb.Search("比基尼", 1)
	tb.Search("nike", 1)

	for {

		_, err := client.LPop(QUEUE_PRODUCTS).Result()
	

		if err == nil {

			//jd := Jd{}

			//jd.ParsePics(r)

			

		}
	
	}
	
} // productJob
