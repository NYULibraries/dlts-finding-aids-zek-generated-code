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
			Origination []struct {
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
		Relatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
				Lb string `xml:"lb"`
			} `xml:"p"`
		} `xml:"relatedmaterial"`
		Arrangement struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"arrangement"`
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
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"bioghist"`
		Custodhist struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"custodhist"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Separatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"separatedmaterial"`
		Processinfo struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"processinfo"`
		Appraisal struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
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
			Geogname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"geogname"`
			Occupation struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
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
							} `xml:"persname"`
						} `xml:"origination"`
						Unitid   string `xml:"unitid"`
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
						Dao struct {
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
					} `xml:"did"`
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
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Separatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"separatedmaterial"`
					Controlaccess struct {
						Text       string `xml:",chardata"`
						Occupation struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"occupation"`
						Geogname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"controlaccess"`
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
							Dao struct {
								Text     string `xml:",chardata"`
								Actuate  string `xml:"actuate,attr"`
								Audience string `xml:"audience,attr"`
								Href     string `xml:"href,attr"`
								Role     string `xml:"role,attr"`
								Show     string `xml:"show,attr"`
								Title    string `xml:"title,attr"`
								Type     string `xml:"type,attr"`
								Daodesc  struct {
									Text string   `xml:",chardata"`
									P    []string `xml:"p"`
								} `xml:"daodesc"`
							} `xml:"dao"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Label    string `xml:"label,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
									Rules  string `xml:"rules,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
						} `xml:"did"`
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
						Odd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"odd"`
						C []struct {
							Text       string `xml:",chardata"`
							ID         string `xml:"id,attr"`
							Level      string `xml:"level,attr"`
							Otherlevel string `xml:"otherlevel,attr"`
							Did        struct {
								Text      string `xml:",chardata"`
								Unittitle struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Title struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"unittitle"`
								Unitid   string `xml:"unitid"`
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
										P    []struct {
											Text string `xml:",chardata"`
											Emph struct {
												Text   string `xml:",chardata"`
												Render string `xml:"render,attr"`
											} `xml:"emph"`
											Title struct {
												Text   string `xml:",chardata"`
												Render string `xml:"render,attr"`
											} `xml:"title"`
										} `xml:"p"`
									} `xml:"daodesc"`
								} `xml:"dao"`
								Physdesc struct {
									Text      string `xml:",chardata"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Altrender string `xml:"altrender,attr"`
									Extent    struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
									} `xml:"extent"`
								} `xml:"physdesc"`
								Origination []struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Persname struct {
										Text   string `xml:",chardata"`
										Source string `xml:"source,attr"`
									} `xml:"persname"`
								} `xml:"origination"`
							} `xml:"did"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    []struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"scopecontent"`
							Odd struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"odd"`
							Separatedmaterial struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"separatedmaterial"`
							Controlaccess struct {
								Text     string `xml:",chardata"`
								Geogname []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"geogname"`
								Subject struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"subject"`
								Persname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"persname"`
								Genreform []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
							} `xml:"controlaccess"`
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
									Container []struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
										ID        string `xml:"id,attr"`
										Label     string `xml:"label,attr"`
										Type      string `xml:"type,attr"`
										Parent    string `xml:"parent,attr"`
									} `xml:"container"`
									Dao struct {
										Text     string `xml:",chardata"`
										Actuate  string `xml:"actuate,attr"`
										Audience string `xml:"audience,attr"`
										Href     string `xml:"href,attr"`
										Role     string `xml:"role,attr"`
										Show     string `xml:"show,attr"`
										Title    string `xml:"title,attr"`
										Type     string `xml:"type,attr"`
										Daodesc  struct {
											Text string   `xml:",chardata"`
											P    []string `xml:"p"`
										} `xml:"daodesc"`
									} `xml:"dao"`
								} `xml:"did"`
								Scopecontent struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    []struct {
										Text string `xml:",chardata"`
										Emph []struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"emph"`
									} `xml:"p"`
								} `xml:"scopecontent"`
							} `xml:"c"`
							Bioghist struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"bioghist"`
							Processinfo struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"processinfo"`
						} `xml:"c"`
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
							Persname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"persname"`
						} `xml:"controlaccess"`
					} `xml:"c"`
					Odd struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"odd"`
				} `xml:"c"`
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

