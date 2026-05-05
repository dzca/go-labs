func main() {
	// 1. Setup Initial Credentials
	serverPath := "://geotab.com" // Start generic, update after Auth
	auth := AuthParams{
		Database: "your_db_name",
		UserName: "your_user@email.com",
		Password: "your_password",
	}

	// 2. Authenticate (Simplified call)
	// Response will provide credentials.SessionId and credentials.Path
	fmt.Println("Authenticating...")
	
	// 3. Start GetFeed Loop
	fromVersion := "" // Start empty to seed the feed
	
	for {
		params := GetFeedParams{
			TypeName:    "ExceptionEvent",
			FromVersion: &fromVersion,
			Credentials: Credentials{
				Database:  auth.Database,
				UserName:  auth.UserName,
				SessionId: "SESSION_TOKEN_FROM_AUTH",
			},
		}

		resp, err := callGeotab(serverPath, "GetFeed", params)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		// Update version for next poll
		fromVersion = resp.Result.ToVersion
		
		// Process speeding events
		for _, event := range resp.Result.Data {
			avgSpeed := calculateAvgSpeed(event.Distance, event.Duration)
			fmt.Printf("Device %s was speeding! Avg Speed: %.2f km/h\n", event.Device.ID, avgSpeed)
			// TODO: Insert into PostgreSQL here
		}
		
		// Sleep for 30 seconds as discussed
		// time.Sleep(30 * time.Second)
	}
}

// Helper to calculate speed based on your sample logic
func calculateAvgSpeed(distKm float64, durationStr string) float64 {
	var h, m, s int
	fmt.Sscanf(durationStr, "%d:%d:%d", &h, &m, &s)
	totalHours := float64(h) + float64(m)/60 + float64(s)/3600
	if totalHours == 0 { return 0 }
	return distKm / totalHours
}
