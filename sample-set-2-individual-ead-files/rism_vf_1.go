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
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
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
			Unitdate []struct {
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
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"scopecontent"`
		Bioghist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
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
		Processinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"processinfo"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"arrangement"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"prefercite"`
		Controlaccess struct {
			Text    string `xml:",chardata"`
			Subject []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"subject"`
			Corpname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"corpname"`
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
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
					Container struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"container"`
					Unitdate struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
				} `xml:"did"`
				C []struct {
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
						Container []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
						} `xml:"container"`
						Unitid   string `xml:"unitid"`
						Unitdate struct {
							Text   string `xml:",chardata"`
							Normal string `xml:"normal,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"unitdate"`
						Origination struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
						} `xml:"origination"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Extent    struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"extent"`
						} `xml:"physdesc"`
						Abstract struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"abstract"`
						Physloc struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"physloc"`
					} `xml:"did"`
					C []struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"id,attr"`
						Level string `xml:"level,attr"`
						Did   struct {
							Text        string `xml:",chardata"`
							Unittitle   string `xml:"unittitle"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Label    string `xml:"label,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
									Rules  string `xml:"rules,attr"`
									Role   string `xml:"role,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
							Unitdate struct {
								Text   string `xml:",chardata"`
								Normal string `xml:"normal,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"unitdate"`
							Physdesc []struct {
								Text      string `xml:",chardata"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Altrender string `xml:"altrender,attr"`
								Extent    []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
								} `xml:"extent"`
							} `xml:"physdesc"`
							Physloc struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"physloc"`
							Container []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Type      string `xml:"type,attr"`
								Parent    string `xml:"parent,attr"`
							} `xml:"container"`
							Unitid       string `xml:"unitid"`
							Langmaterial struct {
								Text     string `xml:",chardata"`
								Language struct {
									Text       string `xml:",chardata"`
									Langcode   string `xml:"langcode,attr"`
									Scriptcode string `xml:"scriptcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
							Abstract struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"abstract"`
						} `xml:"did"`
						Controlaccess struct {
							Text    string `xml:",chardata"`
							Subject []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"subject"`
							Geogname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"geogname"`
							Genreform struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"genreform"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
						} `xml:"controlaccess"`
						C struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Level string `xml:"level,attr"`
							Did   struct {
								Text        string `xml:",chardata"`
								Unittitle   string `xml:"unittitle"`
								Unitid      string `xml:"unitid"`
								Origination struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Persname struct {
										Text   string `xml:",chardata"`
										Source string `xml:"source,attr"`
									} `xml:"persname"`
								} `xml:"origination"`
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
								Physloc struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"physloc"`
								Container struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Type      string `xml:"type,attr"`
								} `xml:"container"`
							} `xml:"did"`
							Controlaccess struct {
								Text     string `xml:",chardata"`
								Geogname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"geogname"`
								Subject struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"subject"`
							} `xml:"controlaccess"`
						} `xml:"c"`
					} `xml:"c"`
					Controlaccess struct {
						Text    string `xml:",chardata"`
						Subject struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Persname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
						Geogname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
					} `xml:"controlaccess"`
				} `xml:"c"`
				Controlaccess struct {
					Text     string `xml:",chardata"`
					Geogname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
				} `xml:"controlaccess"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

