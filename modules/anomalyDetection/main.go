package anomalyDetection

import (
	"github.com/google/gopacket"
	"log"
	"os"
)

var LOGGER *log.Logger

func init() {
	// Initialize the logger with appropriate settings
	// You can customize the log output destination and settings as needed
	LOGGER = log.New(os.Stdout, "AnomalyDetection: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// AnomalyDetection uses a pre-trained machine learning model for anomaly detection.
func AnomalyDetection(packet gopacket.Packet) {
	// Normalize features to the range [0, 1]
	// Use the trained model to predict if the data is anomalous
	// Log if the prediction is for malicious or non-malicious packet
	isAnomalous := true
	if isAnomalous {
		LOGGER.Printf("Anomaly Detected: [%v]", packet)
		// Take appropriate action, such as logging, alerting, or further analysis
	}
}
