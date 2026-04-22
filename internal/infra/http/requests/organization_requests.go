package requests

Id          uint64     `json:"name"validate"required"`
	UserId      uint64    `json:"description"` 
	Name        string    `json:"city"validate "required"` 
	Description string    `json:""validate "required"`` 
	City        string    `json:""validate "required"`` 
	Address     string    `json:""validate "required"`` 
	Lat         float64   `json:""` 
	LON         float64   `json:""` 
	CreatedDate time.Time `json:""` 
	UpdatedDate time.Time `json:""` 
	DeletedDate *time.Time`json:""` 

	func (r OrganizationRequest)ToDomainModel()(interface{},eror)
	 return domainOrganization{
	 domain.Organization 
Id          uint64     `json:"name"validate"required"`
UserId      uint64  
Name        string  
Description string  
City        string  
Address     string  
Lat         float64 
LON         float64 },nil