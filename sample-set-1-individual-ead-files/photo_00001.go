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
				P         struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"p"`
				Address struct {
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
					Text     string `xml:",chardata"`
					Langcode string `xml:"langcode,attr"`
				} `xml:"language"`
			} `xml:"langmaterial"`
		} `xml:"did"`
		Processinfo []struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"processinfo"`
		Arrangement struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"arrangement"`
		Custodhist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"custodhist"`
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
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"scopecontent"`
		Relatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"relatedmaterial"`
		Bioghist struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"bioghist"`
		Controlaccess struct {
			Text     string `xml:",chardata"`
			Geogname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"geogname"`
			Genreform []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"genreform"`
			Subject struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"subject"`
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
							Text       string `xml:",chardata"`
							Langcode   string `xml:"langcode,attr"`
							Scriptcode string `xml:"scriptcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
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
				} `xml:"did"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"scopecontent"`
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
								Text       string `xml:",chardata"`
								Langcode   string `xml:"langcode,attr"`
								Scriptcode string `xml:"scriptcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
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
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Label string `xml:"label,attr"`
						} `xml:"physdesc"`
						Unitid string `xml:"unitid"`
					} `xml:"did"`
					C []struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"id,attr"`
						Level string `xml:"level,attr"`
						Did   struct {
							Text      string `xml:",chardata"`
							Unittitle string `xml:"unittitle"`
							Unitdate  struct {
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
							Physdesc []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
								Emph  struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"physdesc"`
							Unitid string `xml:"unitid"`
							Dao    []struct {
								Text     string `xml:",chardata"`
								Actuate  string `xml:"actuate,attr"`
								Audience string `xml:"audience,attr"`
								Href     string `xml:"href,attr"`
								Role     string `xml:"role,attr"`
								Show     string `xml:"show,attr"`
								Title    string `xml:"title,attr"`
								Type     string `xml:"type,attr"`
								Daodesc  struct {
									Text string `xml:",chardata"`
									P    string `xml:"p"`
								} `xml:"daodesc"`
							} `xml:"dao"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Label    string `xml:"label,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
									Role   string `xml:"role,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
									Rules  string `xml:"rules,attr"`
									Role   string `xml:"role,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
						} `xml:"did"`
						C []struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Level string `xml:"level,attr"`
							Did   struct {
								Text      string `xml:",chardata"`
								Unittitle string `xml:"unittitle"`
								Unitid    string `xml:"unitid"`
								Unitdate  struct {
									Text   string `xml:",chardata"`
									Normal string `xml:"normal,attr"`
									Type   string `xml:"type,attr"`
								} `xml:"unitdate"`
								Physdesc []struct {
									Text   string `xml:",chardata"`
									ID     string `xml:"id,attr"`
									Label  string `xml:"label,attr"`
									Extent string `xml:"extent"`
								} `xml:"physdesc"`
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
								Dao []struct {
									Text     string `xml:",chardata"`
									Actuate  string `xml:"actuate,attr"`
									Audience string `xml:"audience,attr"`
									Href     string `xml:"href,attr"`
									Role     string `xml:"role,attr"`
									Show     string `xml:"show,attr"`
									Title    string `xml:"title,attr"`
									Type     string `xml:"type,attr"`
									Daodesc  struct {
										Text string `xml:",chardata"`
										P    string `xml:"p"`
									} `xml:"daodesc"`
								} `xml:"dao"`
								Origination struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Persname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"persname"`
									Corpname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"corpname"`
								} `xml:"origination"`
							} `xml:"did"`
							Controlaccess struct {
								Text     string `xml:",chardata"`
								Geogname []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"geogname"`
								Genreform struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
								Subject struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"subject"`
							} `xml:"controlaccess"`
							Userestrict struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"userestrict"`
							C []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Level string `xml:"level,attr"`
								Did   struct {
									Text      string `xml:",chardata"`
									Unittitle string `xml:"unittitle"`
									Unitid    string `xml:"unitid"`
									Unitdate  struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
									Physdesc struct {
										Text  string `xml:",chardata"`
										ID    string `xml:"id,attr"`
										Label string `xml:"label,attr"`
									} `xml:"physdesc"`
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
									Dao []struct {
										Text     string `xml:",chardata"`
										Actuate  string `xml:"actuate,attr"`
										Audience string `xml:"audience,attr"`
										Href     string `xml:"href,attr"`
										Role     string `xml:"role,attr"`
										Show     string `xml:"show,attr"`
										Title    string `xml:"title,attr"`
										Type     string `xml:"type,attr"`
										Daodesc  struct {
											Text string `xml:",chardata"`
											P    string `xml:"p"`
										} `xml:"daodesc"`
									} `xml:"dao"`
									Origination struct {
										Text     string `xml:",chardata"`
										Label    string `xml:"label,attr"`
										Corpname struct {
											Text   string `xml:",chardata"`
											Rules  string `xml:"rules,attr"`
											Source string `xml:"source,attr"`
										} `xml:"corpname"`
										Persname struct {
											Text   string `xml:",chardata"`
											Rules  string `xml:"rules,attr"`
											Source string `xml:"source,attr"`
										} `xml:"persname"`
									} `xml:"origination"`
								} `xml:"did"`
								Controlaccess struct {
									Text     string `xml:",chardata"`
									Geogname struct {
										Text   string `xml:",chardata"`
										Source string `xml:"source,attr"`
									} `xml:"geogname"`
								} `xml:"controlaccess"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    string `xml:"p"`
								} `xml:"odd"`
							} `xml:"c"`
							Phystech struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"phystech"`
						} `xml:"c"`
						Controlaccess struct {
							Text     string `xml:",chardata"`
							Geogname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"geogname"`
							Subject []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"subject"`
							Genreform []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"genreform"`
							Persname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
								Role   string `xml:"role,attr"`
							} `xml:"persname"`
							Corpname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"corpname"`
						} `xml:"controlaccess"`
						Odd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"odd"`
						Userestrict struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"userestrict"`
						Accessrestrict struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"accessrestrict"`
					} `xml:"c"`
					Odd struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"odd"`
					Controlaccess struct {
						Text     string `xml:",chardata"`
						Geogname struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
					} `xml:"controlaccess"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"scopecontent"`
				} `xml:"c"`
				Arrangement struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"arrangement"`
				Bioghist struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"bioghist"`
			} `xml:"c"`
		} `xml:"dsc"`
	} `xml:"archdesc"`
} 

