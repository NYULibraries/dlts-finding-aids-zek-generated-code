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
			Text           string `xml:",chardata"`
			Countrycode    string `xml:"countrycode,attr"`
			Mainagencycode string `xml:"mainagencycode,attr"`
			URL            string `xml:"url,attr"`
		} `xml:"eadid"`
		Filedesc struct {
			Text      string `xml:",chardata"`
			Titlestmt struct {
				Text        string `xml:",chardata"`
				Titleproper struct {
					Text string `xml:",chardata"`
					Num  string `xml:"num"`
				} `xml:"titleproper"`
				Author string `xml:"author"`
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
			Langusage string `xml:"langusage"`
			Descrules string `xml:"descrules"`
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
			Unittitle   string `xml:"unittitle"`
			Origination struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Corpname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
				} `xml:"corpname"`
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
		Phystech struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"phystech"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Did   struct {
					Text         string `xml:",chardata"`
					Unittitle    string `xml:"unittitle"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						Language struct {
							Text     string `xml:",chardata"`
							Langcode string `xml:"langcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
				} `xml:"did"`
				C []struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					Level      string `xml:"level,attr"`
					Otherlevel string `xml:"otherlevel,attr"`
					Did        struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"unittitle"`
						Unitid   string `xml:"unitid"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Extent    []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"extent"`
							Dimensions struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
							} `xml:"dimensions"`
							Physfacet struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
							} `xml:"physfacet"`
						} `xml:"physdesc"`
						Unitdate []struct {
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
						Container []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
						} `xml:"container"`
						Origination []struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
						} `xml:"origination"`
					} `xml:"did"`
					Altformavail []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"altformavail"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Phystech struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"phystech"`
					Odd struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"odd"`
					Relatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text   string `xml:",chardata"`
							Extref []struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
								Show    string `xml:"show,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"relatedmaterial"`
					Controlaccess struct {
						Text     string `xml:",chardata"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"persname"`
						Corpname []struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
							Role   string `xml:"role,attr"`
						} `xml:"corpname"`
						Geogname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Genreform struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Title []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"title"`
						Famname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"famname"`
					} `xml:"controlaccess"`
					Custodhist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"custodhist"`
					Otherfindaid struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"otherfindaid"`
					Bioghist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"bioghist"`
					Accessrestrict struct {
						Text        string `xml:",chardata"`
						Legalstatus struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"legalstatus"`
					} `xml:"accessrestrict"`
				} `xml:"c"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

