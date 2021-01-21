type Ead struct {
	XMLName        xml.Name `xml:"ead"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Eadheader      struct {
		Text               string `xml:",chardata"`
		Countryencoding    string `xml:"countryencoding,attr"`
		Dateencoding       string `xml:"dateencoding,attr"`
		Findaidstatus      string `xml:"findaidstatus,attr"`
		Langencoding       string `xml:"langencoding,attr"`
		Repositoryencoding string `xml:"repositoryencoding,attr"`
		Eadid              struct {
			Text        string `xml:",chardata"`
			Countrycode string `xml:"countrycode,attr"`
		} `xml:"eadid"`
		Filedesc struct {
			Text      string `xml:",chardata"`
			Titlestmt struct {
				Text        string `xml:",chardata"`
				Titleproper struct {
					Text string   `xml:",chardata"`
					Lb   []string `xml:"lb"`
					Num  string   `xml:"num"`
				} `xml:"titleproper"`
				Author string `xml:"author"`
			} `xml:"titlestmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher string `xml:"publisher"`
				Address   struct {
					Text        string `xml:",chardata"`
					Addressline string `xml:"addressline"`
				} `xml:"address"`
			} `xml:"publicationstmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"creation"`
			Langusage string `xml:"langusage"`
		} `xml:"profiledesc"`
	} `xml:"eadheader"`
	Archdesc struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
		Did   struct {
			Text       string `xml:",chardata"`
			Repository struct {
				Text     string `xml:",chardata"`
				Corpname string `xml:"corpname"`
			} `xml:"repository"`
			Unittitle string `xml:"unittitle"`
			Unitid    string `xml:"unitid"`
			Physdesc  struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				Extent    []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
			} `xml:"physdesc"`
			Unitdate []struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Container []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Label string `xml:"label,attr"`
				Type  string `xml:"type,attr"`
			} `xml:"container"`
		} `xml:"did"`
		Bioghist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"bioghist"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"arrangement"`
		Accessrestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accessrestrict"`
		Custodhist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"custodhist"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"scopecontent"`
		Controlaccess struct {
			Text      string `xml:",chardata"`
			Genreform []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"genreform"`
			Geogname struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"geogname"`
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Did   struct {
					Text      string `xml:",chardata"`
					Unittitle string `xml:"unittitle"`
					Unitid    struct {
						Text       string `xml:",chardata"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Corpname struct {
							Text           string `xml:",chardata"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"corpname"`
						Persname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"origination"`
					Unitdate struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
					Container []struct {
						Text   string `xml:",chardata"`
						ID     string `xml:"id,attr"`
						Label  string `xml:"label,attr"`
						Type   string `xml:"type,attr"`
						Parent string `xml:"parent,attr"`
					} `xml:"container"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						ID        string `xml:"id,attr"`
						Altrender string `xml:"altrender,attr"`
						Extent    struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
						Dimensions struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"dimensions"`
					} `xml:"physdesc"`
				} `xml:"did"`
				C []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle string `xml:"unittitle"`
						Unitid    struct {
							Text       string `xml:",chardata"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
						Origination struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"persname"`
						} `xml:"origination"`
						Unitdate  string `xml:"unitdate"`
						Container []struct {
							Text   string `xml:",chardata"`
							ID     string `xml:"id,attr"`
							Label  string `xml:"label,attr"`
							Type   string `xml:"type,attr"`
							Parent string `xml:"parent,attr"`
						} `xml:"container"`
					} `xml:"did"`
					Dao []struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Role    string `xml:"role,attr"`
						Show    string `xml:"show,attr"`
						Title   string `xml:"title,attr"`
						Type    string `xml:"type,attr"`
						Daodesc struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Controlaccess struct {
						Text      string `xml:",chardata"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Geogname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Occupation struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"occupation"`
					} `xml:"controlaccess"`
				} `xml:"c"`
				Dao []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Role    string `xml:"role,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Daodesc struct {
						Text string   `xml:",chardata"`
						P    []string `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Controlaccess struct {
					Text      string `xml:",chardata"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Subject []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"subject"`
					Geogname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
					Occupation []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"occupation"`
				} `xml:"controlaccess"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

