provider "google" {
    credentials = "$file("CREDENTIALS_FILE.json")"
    project = "MeetingsIBMHackathon"
    region = "eu-west1"
}