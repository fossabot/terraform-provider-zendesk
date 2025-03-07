# ticket_fields.tf
#   Ticket Field example
#
# API reference:
#   https://developer.zendesk.com/rest_api/docs/support/ticket_fields

resource "zendesk_ticket_field" "checkbox-field" {
  title = "Checkbox Field"
  type = "checkbox"
}

resource "zendesk_ticket_field" "date-field" {
  title = "Date Field"
  type = "date"
}

resource "zendesk_ticket_field" "decimal-field" {
  title = "Decimal Field"
  type = "decimal"
}

resource "zendesk_ticket_field" "integer-field" {
  title = "Integer Field"
  type = "integer"
}

resource "zendesk_ticket_field" "regexp-field" {
  title = "Regexp Field"
  type = "regexp"
  regexp_for_validation = "^[0-9]+-[0-9]+-[0-9]+$"
}

resource "zendesk_ticket_field" "tagger-field" {
  title = "Tagger Field"
  type = "tagger"

  custom_field_option {
    name  = "Option 1"
    value = "opt1"
  }

  custom_field_option {
    name  = "Option 2"
    value = "opt2"
  }
}

resource "zendesk_ticket_field" "text-field" {
  title = "Text Field"
  type = "text"
}

resource "zendesk_ticket_field" "textarea-field" {
  title = "Textarea Field"
  type = "textarea"
}

variable "ASSIGNEE_TICKET_FIELD_ID" { type = "string" }
data "zendesk_ticket_field" "assignee" {
  id = "${var.ASSIGNEE_TICKET_FIELD_ID}"
}

variable "GROUP_TICKET_FIELD_ID" { type = "string" }
data "zendesk_ticket_field" "group" {
  id = "${var.GROUP_TICKET_FIELD_ID}"
}

variable "STATUS_TICKET_FIELD_ID" { type = "string" }
data "zendesk_ticket_field" "status" {
  id = "${var.STATUS_TICKET_FIELD_ID}"
}

variable "SUBJECT_TICKET_FIELD_ID" { type = "string" }
data "zendesk_ticket_field" "subject" {
  id = "${var.SUBJECT_TICKET_FIELD_ID}"
}

variable "DESCRIPTION_TICKET_FIELD_ID" { type = "string" }
data "zendesk_ticket_field" "description" {
  id = "${var.DESCRIPTION_TICKET_FIELD_ID}"
}