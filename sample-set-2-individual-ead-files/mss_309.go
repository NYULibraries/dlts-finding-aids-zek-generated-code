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
		Langencoding       string `xml:"langencoding,attr"`
		Repositoryencoding string `xml:"repositoryencoding,attr"`
		Eadid              struct {
			Text           string `xml:",chardata"`
			Countrycode    string `xml:"countrycode,attr"`
			Mainagencycode string `xml:"mainagencycode,attr"`
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
			} `xml:"titlestmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher string `xml:"publisher"`
			} `xml:"publicationstmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"creation"`
			Langusage struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text           string `xml:",chardata"`
					Langcode       string `xml:"langcode,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"language"`
			} `xml:"langusage"`
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
			Unittitle struct {
				Text string `xml:",chardata"`
				Lb   string `xml:"lb"`
			} `xml:"unittitle"`
			Origination struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Persname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"persname"`
			} `xml:"origination"`
			Unitid   string `xml:"unitid"`
			Physdesc struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				Extent    struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
			} `xml:"physdesc"`
			Unitdate struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Langmaterial struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
			} `xml:"langmaterial"`
		} `xml:"did"`
		Bioghist []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"bioghist"`
		Accessrestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accessrestrict"`
		Userestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"userestrict"`
		Controlaccess struct {
			Text    string `xml:",chardata"`
			Subject []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"subject"`
			Genreform []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"genreform"`
			Occupation struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
			Persname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Rules  string `xml:"rules,attr"`
			} `xml:"persname"`
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Did   struct {
					Text        string `xml:",chardata"`
					Unittitle   string `xml:"unittitle"`
					Origination struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
				} `xml:"did"`
				Scopecontent []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"scopecontent"`
				Controlaccess struct {
					Text    string `xml:",chardata"`
					Subject []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"subject"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
						Rules  string `xml:"rules,attr"`
					} `xml:"corpname"`
					Persname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
						Rules  string `xml:"rules,attr"`
					} `xml:"persname"`
					Geogname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Occupation []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"occupation"`
				} `xml:"controlaccess"`
				C []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle string `xml:"unittitle"`
						Unitdate  string `xml:"unitdate"`
						Dao       struct {
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
					} `xml:"did"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"scopecontent"`
				} `xml:"c"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

