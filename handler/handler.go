

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.Gethostname()

	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	//fmt.Fprintf(w, hostname, IP)
	response := map[string]string{
		"hostname": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}