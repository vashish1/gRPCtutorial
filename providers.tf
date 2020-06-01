provider "google" {
  credentials = "${file("<file>.json")}"
  project     = "${var.gcloud-project}"
  region      = "${var.gcloud-region}"
}