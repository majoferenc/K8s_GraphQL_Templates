resource "google_dns_record_set" "my_awesome_service" {
 project = "MeetingsIBMHackathon"
 managed_zone = "meetingsibmhackathon"
 name = "meetingsibmhackathon.com"
 type = "A"
 ttl = 300
 rrdatas = ["${google_compute_global_address.my_awesome_service.address}"]
}
