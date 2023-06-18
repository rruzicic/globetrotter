package models

const (
	HostRole  string = "HOST"
	GuestRole string = "GUEST"
)

type User struct {
	Model                `bson:",inline"`
	FirstName            string `json:"firstName" bson:"first_name"`
	LastName             string `json:"lastName" bson:"last_name"`
	EMail                string `json:"email" bson:"email"`
	Password             string `json:"password" bson:"password"`
	Role                 string `json:"role" bson:"role"`
	SuperHost            bool   `json:"superHost" bson:"super_host"`
	CancellationsCounter int    `json:"cancellationsCounter" bson:"cancellations_counter"`
	RatingStatus
	Address
	ReservationStatus
}
