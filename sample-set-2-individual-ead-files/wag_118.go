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
			Text string `xml:",chardata"`
			URL  string `xml:"url,attr"`
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
				Address   struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text   string `xml:",chardata"`
						Extptr struct {
							Text  string `xml:",chardata"`
							Href  string `xml:"href,attr"`
							Show  string `xml:"show,attr"`
							Title string `xml:"title,attr"`
							Type  string `xml:"type,attr"`
						} `xml:"extptr"`
					} `xml:"addressline"`
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
			Descrules string `xml:"descrules"`
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Item string `xml:"item"`
			} `xml:"change"`
		} `xml:"revisiondesc"`
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
			Origination []struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Corpname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"corpname"`
				Persname struct {
					Text   string `xml:",chardata"`
					Role   string `xml:"role,attr"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
				} `xml:"persname"`
			} `xml:"origination"`
			Unitid   string `xml:"unitid"`
			Physdesc []struct {
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
			Physloc struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"physloc"`
			Abstract struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"abstract"`
			Langmaterial struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"langmaterial"`
		} `xml:"did"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"arrangement"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"scopecontent"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"prefercite"`
		Acqinfo struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"acqinfo"`
		Processinfo struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"processinfo"`
		Phystech []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"phystech"`
		Accessrestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accessrestrict"`
		Bioghist struct {
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
		} `xml:"bioghist"`
		Userestrict struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"userestrict"`
		Custodhist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"custodhist"`
		Accruals struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accruals"`
		Appraisal struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"appraisal"`
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
			Occupation struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
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
					Container []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"container"`
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
					P    string `xml:"p"`
				} `xml:"scopecontent"`
				Phystech struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"phystech"`
				Accruals struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"accruals"`
				Prefercite struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"prefercite"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

