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
					Text string   `xml:",chardata"`
					Lb   []string `xml:"lb"`
					Num  string   `xml:"num"`
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
			Origination struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Corpname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
				} `xml:"corpname"`
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
			Langmaterial struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
			} `xml:"langmaterial"`
		} `xml:"did"`
		Accessrestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accessrestrict"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"prefercite"`
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
		Relatedmaterial struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"relatedmaterial"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
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
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"bioghist"`
		Separatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"separatedmaterial"`
		Phystech struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"phystech"`
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
			Title struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"title"`
			Subject []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"subject"`
			Corpname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"corpname"`
			Persname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"persname"`
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
				} `xml:"did"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
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
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"container"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							Extent    struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"extent"`
							Dimensions struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"dimensions"`
							Physfacet struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"physfacet"`
						} `xml:"physdesc"`
					} `xml:"did"`
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
							Physdesc []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								Extent    []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
									Lb        string `xml:"lb"`
								} `xml:"extent"`
								Physfacet struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"physfacet"`
								Dimensions struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"dimensions"`
							} `xml:"physdesc"`
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
								Parent    string `xml:"parent,attr"`
							} `xml:"container"`
							Unitid string `xml:"unitid"`
						} `xml:"did"`
						Phystech struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"phystech"`
						Altformavail struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"altformavail"`
						Scopecontent struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"scopecontent"`
						Separatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"separatedmaterial"`
					} `xml:"c"`
					Phystech struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"phystech"`
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
					Altformavail struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"altformavail"`
				} `xml:"c"`
				Accessrestrict struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"accessrestrict"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

