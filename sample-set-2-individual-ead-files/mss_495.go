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
				P         struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"p"`
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
				Persname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
				} `xml:"persname"`
			} `xml:"origination"`
			Unitid   string `xml:"unitid"`
			Physdesc struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				Extent    []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
			} `xml:"physdesc"`
			Unitdate struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Abstract struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"abstract"`
			Physloc struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"physloc"`
			Langmaterial struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"langmaterial"`
		} `xml:"did"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"scopecontent"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"arrangement"`
		Bioghist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
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
		Acqinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"acqinfo"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"prefercite"`
		Processinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"processinfo"`
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
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				Level      string `xml:"level,attr"`
				Otherlevel string `xml:"otherlevel,attr"`
				Did        struct {
					Text      string `xml:",chardata"`
					Unittitle string `xml:"unittitle"`
					Unitdate  struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
					Container struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"container"`
					Origination struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Persname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"origination"`
					Physdesc struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
					} `xml:"physdesc"`
					Dao struct {
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
					P    struct {
						Text  string   `xml:",chardata"`
						Title []string `xml:"title"`
					} `xml:"p"`
				} `xml:"scopecontent"`
				Arrangement struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"arrangement"`
				C []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"unittitle"`
						Unitdate struct {
							Text   string `xml:",chardata"`
							Normal string `xml:"normal,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"unitdate"`
						Container []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
						} `xml:"container"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							Language struct {
								Text     string `xml:",chardata"`
								Langcode string `xml:"langcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
					} `xml:"did"`
				} `xml:"c"`
				Accruals struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"accruals"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

