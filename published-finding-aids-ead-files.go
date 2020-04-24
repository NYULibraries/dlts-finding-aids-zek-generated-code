type Ead struct {
	XMLName        xml.Name `xml:"ead"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Eadheader      struct {
		Text               string `xml:",chardata"`
		Countryencoding    string `xml:"countryencoding,attr"`
		Dateencoding       string `xml:"dateencoding,attr"`
		Langencoding       string `xml:"langencoding,attr"`
		Repositoryencoding string `xml:"repositoryencoding,attr"`
		Findaidstatus      string `xml:"findaidstatus,attr"`
		Audience           string `xml:"audience,attr"`
		ID                 string `xml:"id,attr"`
		Eadid              struct {
			Text           string `xml:",chardata"`
			Countrycode    string `xml:"countrycode,attr"`
			URL            string `xml:"url,attr"`
			Mainagencycode string `xml:"mainagencycode,attr"`
			Identifier     string `xml:"identifier,attr"`
		} `xml:"eadid"`
		Filedesc struct {
			Text      string `xml:",chardata"`
			Titlestmt struct {
				Text        string `xml:",chardata"`
				Titleproper []struct {
					Text           string   `xml:",chardata"`
					Type           string   `xml:"type,attr"`
					Encodinganalog string   `xml:"encodinganalog,attr"`
					Num            []string `xml:"num"`
					Lb             []string `xml:"lb"`
					Date           []struct {
						Text     string `xml:",chardata"`
						Calendar string `xml:"calendar,attr"`
						Normal   string `xml:"normal,attr"`
						Era      string `xml:"era,attr"`
						Type     string `xml:"type,attr"`
						Lb       string `xml:"lb"`
					} `xml:"date"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"titleproper"`
				Author struct {
					Text           string `xml:",chardata"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Lb             string `xml:"lb"`
				} `xml:"author"`
				Subtitle string `xml:"subtitle"`
				Sponsor  struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"sponsor"`
			} `xml:"titlestmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher struct {
					Text           string `xml:",chardata"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"publisher"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text   string   `xml:",chardata"`
						Lb     []string `xml:"lb"`
						Extptr struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extptr"`
					} `xml:"addressline"`
				} `xml:"address"`
				Date string `xml:"date"`
				P    struct {
					Text string `xml:",chardata"`
					Date struct {
						Text           string `xml:",chardata"`
						Encodinganalog string `xml:"encodinganalog,attr"`
						Normal         string `xml:"normal,attr"`
					} `xml:"date"`
					Address struct {
						Text        string `xml:",chardata"`
						Addressline string `xml:"addressline"`
					} `xml:"address"`
					Extref struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"p"`
			} `xml:"publicationstmt"`
			Editionstmt struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"`
					Date struct {
						Text     string `xml:",chardata"`
						Calendar string `xml:"calendar,attr"`
						Era      string `xml:"era,attr"`
					} `xml:"date"`
				} `xml:"p"`
			} `xml:"editionstmt"`
			Notestmt struct {
				Text string `xml:",chardata"`
				Note struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"note"`
			} `xml:"notestmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text           string `xml:",chardata"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Date           struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"date"`
			} `xml:"creation"`
			Langusage struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text           string `xml:",chardata"`
					Langcode       string `xml:"langcode,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Scriptcode     string `xml:"scriptcode,attr"`
				} `xml:"language"`
				Lb string `xml:"lb"`
			} `xml:"langusage"`
			Descrules string `xml:"descrules"`
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Item struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"change"`
		} `xml:"revisiondesc"`
	} `xml:"eadheader"`
	Archdesc struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
		Type  string `xml:"type,attr"`
		Did   struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Repository struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Corpname       string `xml:"corpname"`
				Address        struct {
					Text        string `xml:",chardata"`
					Addressline string `xml:"addressline"`
				} `xml:"address"`
			} `xml:"repository"`
			Unittitle struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Emph           []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Unitdate struct {
					Text           string `xml:",chardata"`
					Normal         string `xml:"normal,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Type           string `xml:"type,attr"`
				} `xml:"unitdate"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"unittitle"`
			Origination []struct {
				Text           string `xml:",chardata"`
				Audience       string `xml:"audience,attr"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Persname       struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Rules          string `xml:"rules,attr"`
					Source         string `xml:"source,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"persname"`
				Corpname struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Source         string `xml:"source,attr"`
					Rules          string `xml:"rules,attr"`
					Authfilenumber string `xml:"authfilenumber,attr"`
				} `xml:"corpname"`
				Famname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"famname"`
			} `xml:"origination"`
			Unitid []struct {
				Text           string `xml:",chardata"`
				Label          string `xml:"label,attr"`
				Countrycode    string `xml:"countrycode,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Repositorycode string `xml:"repositorycode,attr"`
				Audience       string `xml:"audience,attr"`
				Identifier     string `xml:"identifier,attr"`
				Type           string `xml:"type,attr"`
			} `xml:"unitid"`
			Physdesc []struct {
				Text           string `xml:",chardata"`
				Altrender      string `xml:"altrender,attr"`
				ID             string `xml:"id,attr"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Audience       string `xml:"audience,attr"`
				Extent         []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
				Physfacet string `xml:"physfacet"`
			} `xml:"physdesc"`
			Unitdate []struct {
				Text      string `xml:",chardata"`
				Normal    string `xml:"normal,attr"`
				Type      string `xml:"type,attr"`
				Certainty string `xml:"certainty,attr"`
			} `xml:"unitdate"`
			Container []struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Type      string `xml:"type,attr"`
				Altrender string `xml:"altrender,attr"`
				Parent    string `xml:"parent,attr"`
			} `xml:"container"`
			Abstract []struct {
				Text           string `xml:",chardata"`
				ID             string `xml:"id,attr"`
				Label          string `xml:"label,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Emph           []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Title  string `xml:"title,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
				Lb []string `xml:"lb"`
			} `xml:"abstract"`
			Materialspec struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Label string `xml:"label,attr"`
			} `xml:"materialspec"`
			Langmaterial []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Label    string `xml:"label,attr"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"langmaterial"`
			Physloc []struct {
				Text           string `xml:",chardata"`
				ID             string `xml:"id,attr"`
				Label          string `xml:"label,attr"`
				Audience       string `xml:"audience,attr"`
				Encodinganalog string `xml:"encodinganalog,attr"`
				Extref         struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Lb   []string `xml:"lb"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"physloc"`
			Head string `xml:"head"`
		} `xml:"did"`
		Scopecontent []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text      string   `xml:",chardata"`
				Genreform []string `xml:"genreform"`
				Date      struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Emph []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Emph      struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
				} `xml:"emph"`
				List []struct {
					Text string `xml:",chardata"`
					Head string `xml:"head"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Lb   []string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
					} `xml:"item"`
				} `xml:"list"`
				Corpname []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"corpname"`
				Name  []string `xml:"name"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Ender  string `xml:"ender,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				Occupation string `xml:"occupation"`
				Archref    []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Actuate string `xml:"actuate,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Geogname string `xml:"geogname"`
				I        string `xml:"i"`
				Persname string `xml:"persname"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Item       []struct {
					Text  string   `xml:",chardata"`
					Lb    []string `xml:"lb"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Archref []struct {
						Text   string `xml:",chardata"`
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"archref"`
					Corpname []string `xml:"corpname"`
				} `xml:"item"`
				Head    string `xml:"head"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  string `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
			Chronlist struct {
				Text      string `xml:",chardata"`
				Head      string `xml:"head"`
				Chronitem []struct {
					Text  string `xml:",chardata"`
					Date  string `xml:"date"`
					Event string `xml:"event"`
				} `xml:"chronitem"`
			} `xml:"chronlist"`
			Extref struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
		} `xml:"scopecontent"`
		Controlaccess struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			Genreform []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"genreform"`
			Geogname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"geogname"`
			Subject []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"subject"`
			Persname []struct {
				Text           string `xml:",chardata"`
				Audience       string `xml:"audience,attr"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"persname"`
			Occupation []struct {
				Text           string `xml:",chardata"`
				Authfilenumber string `xml:"authfilenumber,attr"`
				Source         string `xml:"source,attr"`
			} `xml:"occupation"`
			Corpname []struct {
				Text           string `xml:",chardata"`
				Role           string `xml:"role,attr"`
				Rules          string `xml:"rules,attr"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
				Audience       string `xml:"audience,attr"`
			} `xml:"corpname"`
			Function struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"function"`
			Famname []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Rules          string `xml:"rules,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
				Role           string `xml:"role,attr"`
				Audience       string `xml:"audience,attr"`
			} `xml:"famname"`
			Head          string `xml:"head"`
			Controlaccess []struct {
				Text     string `xml:",chardata"`
				Head     string `xml:"head"`
				Persname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"persname"`
				Famname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"famname"`
				Corpname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
					Role           string `xml:"role,attr"`
				} `xml:"corpname"`
				Subject []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"subject"`
				Geogname []struct {
					Text           string `xml:",chardata"`
					Role           string `xml:"role,attr"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"geogname"`
				Genreform []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"genreform"`
			} `xml:"controlaccess"`
			Title []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"title"`
		} `xml:"controlaccess"`
		Dsc struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Type string `xml:"type,attr"`
			C01  []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Level string `xml:"level,attr"`
				Did   struct {
					Text      string `xml:",chardata"`
					Unittitle struct {
						Text           string `xml:",chardata"`
						Encodinganalog string `xml:"encodinganalog,attr"`
						Emph           []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"unittitle"`
					Unitid []struct {
						Text       string `xml:",chardata"`
						Audience   string `xml:"audience,attr"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Audience string `xml:"audience,attr"`
						Persname struct {
							Text           string `xml:",chardata"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"persname"`
						Corpname struct {
							Text           string `xml:",chardata"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
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
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"dimensions"`
					} `xml:"physdesc"`
					Unitdate []struct {
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
					Langmaterial struct {
						Text     string `xml:",chardata"`
						Language struct {
							Text     string `xml:",chardata"`
							Langcode string `xml:"langcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
				} `xml:"did"`
				Dao []struct {
					Text     string `xml:",chardata"`
					Actuate  string `xml:"actuate,attr"`
					Href     string `xml:"href,attr"`
					Role     string `xml:"role,attr"`
					Show     string `xml:"show,attr"`
					Title    string `xml:"title,attr"`
					Type     string `xml:"type,attr"`
					Audience string `xml:"audience,attr"`
					Daodesc  struct {
						Text string `xml:",chardata"`
						P    string `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Controlaccess struct {
					Text     string `xml:",chardata"`
					Geogname []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"geogname"`
					Subject []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"subject"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Occupation []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"occupation"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Role   string `xml:"role,attr"`
						Rules  string `xml:"rules,attr"`
						Source string `xml:"source,attr"`
					} `xml:"corpname"`
					Title struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"title"`
					Persname struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"persname"`
				} `xml:"controlaccess"`
				C02 []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text  string   `xml:",chardata"`
							Lb    []string `xml:"lb"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"emph"`
						} `xml:"unittitle"`
						Unitid []struct {
							Text       string `xml:",chardata"`
							Audience   string `xml:"audience,attr"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
						Origination []struct {
							Text     string `xml:",chardata"`
							Audience string `xml:"audience,attr"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text           string `xml:",chardata"`
								Role           string `xml:"role,attr"`
								Rules          string `xml:"rules,attr"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text           string `xml:",chardata"`
								Role           string `xml:"role,attr"`
								Source         string `xml:"source,attr"`
								Rules          string `xml:"rules,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"persname"`
						} `xml:"origination"`
						Unitdate []struct {
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
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Extent    []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								Unit      string `xml:"unit,attr"`
							} `xml:"extent"`
							Dimensions struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"dimensions"`
						} `xml:"physdesc"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							Language struct {
								Text     string `xml:",chardata"`
								Langcode string `xml:"langcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
					} `xml:"did"`
					Dao []struct {
						Text     string `xml:",chardata"`
						Actuate  string `xml:"actuate,attr"`
						Href     string `xml:"href,attr"`
						Role     string `xml:"role,attr"`
						Show     string `xml:"show,attr"`
						Title    string `xml:"title,attr"`
						Type     string `xml:"type,attr"`
						Audience string `xml:"audience,attr"`
						Daodesc  struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Controlaccess struct {
						Text    string `xml:",chardata"`
						Subject []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"subject"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Occupation []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"occupation"`
						Geogname []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"geogname"`
						Title struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"title"`
						Corpname []struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"controlaccess"`
					Arrangement struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"arrangement"`
					C03 []struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"id,attr"`
						Level string `xml:"level,attr"`
						Did   struct {
							Text      string `xml:",chardata"`
							Unittitle struct {
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
							} `xml:"unittitle"`
							Unitid []struct {
								Text       string `xml:",chardata"`
								Audience   string `xml:"audience,attr"`
								Identifier string `xml:"identifier,attr"`
								Type       string `xml:"type,attr"`
							} `xml:"unitid"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Audience string `xml:"audience,attr"`
								Label    string `xml:"label,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Role   string `xml:"role,attr"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Role   string `xml:"role,attr"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
							Unitdate []struct {
								Text     string `xml:",chardata"`
								Normal   string `xml:"normal,attr"`
								Type     string `xml:"type,attr"`
								Era      string `xml:"era,attr"`
								Calendar string `xml:"calendar,attr"`
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
							Langmaterial struct {
								Text     string `xml:",chardata"`
								Language struct {
									Text     string `xml:",chardata"`
									Langcode string `xml:"langcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
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
								Text           string `xml:",chardata"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"subject"`
							Geogname []struct {
								Text           string `xml:",chardata"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"geogname"`
							Occupation []struct {
								Text           string `xml:",chardata"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"occupation"`
						} `xml:"controlaccess"`
						C04 []struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Level string `xml:"level,attr"`
							Did   struct {
								Text      string `xml:",chardata"`
								Unittitle struct {
									Text  string `xml:",chardata"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
									Lb   []string `xml:"lb"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"unittitle"`
								Container []struct {
									Text   string `xml:",chardata"`
									ID     string `xml:"id,attr"`
									Type   string `xml:"type,attr"`
									Label  string `xml:"label,attr"`
									Parent string `xml:"parent,attr"`
								} `xml:"container"`
								Unitdate struct {
									Text   string `xml:",chardata"`
									Normal string `xml:"normal,attr"`
									Type   string `xml:"type,attr"`
								} `xml:"unitdate"`
							} `xml:"did"`
							Odd struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    struct {
									Text  string   `xml:",chardata"`
									Lb    []string `xml:"lb"`
									Title []struct {
										Text   string   `xml:",chardata"`
										Render string   `xml:"render,attr"`
										Lb     []string `xml:"lb"`
									} `xml:"title"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"odd"`
							C05 []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Level string `xml:"level,attr"`
								Did   struct {
									Text      string `xml:",chardata"`
									Unittitle struct {
										Text  string `xml:",chardata"`
										Lb    string `xml:"lb"`
										Title struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"title"`
									} `xml:"unittitle"`
									Container []struct {
										Text   string `xml:",chardata"`
										ID     string `xml:"id,attr"`
										Type   string `xml:"type,attr"`
										Label  string `xml:"label,attr"`
										Parent string `xml:"parent,attr"`
									} `xml:"container"`
									Unitdate struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
								} `xml:"did"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    struct {
										Text  string `xml:",chardata"`
										Title struct {
											Text   string   `xml:",chardata"`
											Render string   `xml:"render,attr"`
											Lb     []string `xml:"lb"`
										} `xml:"title"`
										Lb []string `xml:"lb"`
									} `xml:"p"`
								} `xml:"odd"`
							} `xml:"c05"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"scopecontent"`
						} `xml:"c04"`
						Odd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text  string   `xml:",chardata"`
								Lb    []string `xml:"lb"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"odd"`
						Scopecontent struct {
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
						} `xml:"scopecontent"`
						Phystech struct {
							Text string   `xml:",chardata"`
							P    []string `xml:"p"`
						} `xml:"phystech"`
					} `xml:"c03"`
					Altformavail struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"altformavail"`
					Custodhist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"custodhist"`
					Fileplan struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"fileplan"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text  string `xml:",chardata"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Accessrestrict struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accessrestrict"`
					Phystech struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"phystech"`
					Relatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Show    string `xml:"show,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"relatedmaterial"`
					Odd []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"odd"`
				} `xml:"c02"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Emph   []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Lb []string `xml:"lb"`
					} `xml:"p"`
					Head string `xml:"head"`
				} `xml:"scopecontent"`
				Arrangement struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"arrangement"`
				Bioghist struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"bioghist"`
				Odd struct {
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
				} `xml:"odd"`
			} `xml:"c01"`
			C []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				Level      string `xml:"level,attr"`
				Otherlevel string `xml:"otherlevel,attr"`
				Audience   string `xml:"audience,attr"`
				Did        struct {
					Text      string `xml:",chardata"`
					Unittitle struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text      string `xml:",chardata"`
							Render    string `xml:"render,attr"`
							AttrTitle string `xml:"title,attr"`
							Title     struct {
								Text   string `xml:",chardata"`
								Type   string `xml:"type,attr"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"emph"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Lb   []string `xml:"lb"`
						Name struct {
							Text string `xml:",chardata"`
							Role string `xml:"role,attr"`
						} `xml:"name"`
						Corpname string `xml:"corpname"`
						Persname string `xml:"persname"`
					} `xml:"unittitle"`
					Unitid []struct {
						Text       string `xml:",chardata"`
						Audience   string `xml:"audience,attr"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Audience string `xml:"audience,attr"`
						Persname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
							Role   string `xml:"role,attr"`
						} `xml:"persname"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Audience  string `xml:"audience,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
						Physfacet struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"physfacet"`
						Dimensions struct {
							Text     string `xml:",chardata"`
							Audience string `xml:"audience,attr"`
						} `xml:"dimensions"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"physdesc"`
					Unitdate []struct {
						Text      string `xml:",chardata"`
						Normal    string `xml:"normal,attr"`
						Type      string `xml:"type,attr"`
						Certainty string `xml:"certainty,attr"`
					} `xml:"unitdate"`
					Container []struct {
						Text      string `xml:",chardata"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
						Altrender string `xml:"altrender,attr"`
						Parent    string `xml:"parent,attr"`
					} `xml:"container"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Language struct {
							Text     string `xml:",chardata"`
							Langcode string `xml:"langcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
					Physloc struct {
						Text    string `xml:",chardata"`
						ID      string `xml:"id,attr"`
						Label   string `xml:"label,attr"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
					} `xml:"physloc"`
					Abstract struct {
						Text  string   `xml:",chardata"`
						ID    string   `xml:"id,attr"`
						Label string   `xml:"label,attr"`
						Lb    []string `xml:"lb"`
					} `xml:"abstract"`
					Materialspec struct {
						Text  string `xml:",chardata"`
						ID    string `xml:"id,attr"`
						Label string `xml:"label,attr"`
					} `xml:"materialspec"`
				} `xml:"did"`
				Dao []struct {
					Text     string `xml:",chardata"`
					Actuate  string `xml:"actuate,attr"`
					Href     string `xml:"href,attr"`
					Role     string `xml:"role,attr"`
					Show     string `xml:"show,attr"`
					Title    string `xml:"title,attr"`
					Type     string `xml:"type,attr"`
					Audience string `xml:"audience,attr"`
					Daodesc  struct {
						Text string `xml:",chardata"`
						P    struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Controlaccess struct {
					Text     string `xml:",chardata"`
					Geogname []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"geogname"`
					Subject []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"subject"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Occupation []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"occupation"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Role   string `xml:"role,attr"`
						Rules  string `xml:"rules,attr"`
						Source string `xml:"source,attr"`
					} `xml:"corpname"`
					Title []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"title"`
					Persname []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Role           string `xml:"role,attr"`
						Rules          string `xml:"rules,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"persname"`
					Famname struct {
						Text   string `xml:",chardata"`
						Rules  string `xml:"rules,attr"`
						Source string `xml:"source,attr"`
					} `xml:"famname"`
					Function struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"function"`
				} `xml:"controlaccess"`
				C []struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					Level      string `xml:"level,attr"`
					Otherlevel string `xml:"otherlevel,attr"`
					Audience   string `xml:"audience,attr"`
					Did        struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Title  string `xml:"title,attr"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Lb string `xml:"lb"`
							} `xml:"emph"`
							Title []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Type   string   `xml:"type,attr"`
								Lb     []string `xml:"lb"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"title"`
							Lb     []string `xml:"lb"`
							Extref struct {
								Text     string `xml:",chardata"`
								Linktype string `xml:"linktype,attr"`
								Actuate  string `xml:"actuate,attr"`
								Show     string `xml:"show,attr"`
								Href     string `xml:"href,attr"`
							} `xml:"extref"`
							Name struct {
								Text string `xml:",chardata"`
								Role string `xml:"role,attr"`
							} `xml:"name"`
							I        []string `xml:"i"`
							Persname string   `xml:"persname"`
							Corpname string   `xml:"corpname"`
						} `xml:"unittitle"`
						Unitid []struct {
							Text       string `xml:",chardata"`
							Audience   string `xml:"audience,attr"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
						Origination []struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Audience string `xml:"audience,attr"`
							Persname struct {
								Text           string `xml:",chardata"`
								Role           string `xml:"role,attr"`
								Rules          string `xml:"rules,attr"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"persname"`
							Corpname struct {
								Text           string `xml:",chardata"`
								Role           string `xml:"role,attr"`
								Rules          string `xml:"rules,attr"`
								Source         string `xml:"source,attr"`
								Authfilenumber string `xml:"authfilenumber,attr"`
							} `xml:"corpname"`
							Famname struct {
								Text   string `xml:",chardata"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"famname"`
						} `xml:"origination"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Audience  string `xml:"audience,attr"`
							Extent    []struct {
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
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"physfacet"`
						} `xml:"physdesc"`
						Unitdate []struct {
							Text      string `xml:",chardata"`
							Normal    string `xml:"normal,attr"`
							Type      string `xml:"type,attr"`
							Certainty string `xml:"certainty,attr"`
						} `xml:"unitdate"`
						Container []struct {
							Text      string `xml:",chardata"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"container"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Label    string `xml:"label,attr"`
							Language struct {
								Text     string `xml:",chardata"`
								Langcode string `xml:"langcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
						Physloc struct {
							Text  string `xml:",chardata"`
							ID    string `xml:"id,attr"`
							Label string `xml:"label,attr"`
						} `xml:"physloc"`
						Abstract struct {
							Text  string   `xml:",chardata"`
							ID    string   `xml:"id,attr"`
							Label string   `xml:"label,attr"`
							Lb    []string `xml:"lb"`
						} `xml:"abstract"`
						Materialspec struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"materialspec"`
					} `xml:"did"`
					C []struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						Level      string `xml:"level,attr"`
						Otherlevel string `xml:"otherlevel,attr"`
						Audience   string `xml:"audience,attr"`
						Did        struct {
							Text      string `xml:",chardata"`
							Unittitle struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
									Emph   string `xml:"emph"`
									Title  struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"emph"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
									Type   string `xml:"type,attr"`
									Rende  string `xml:"rende,attr"`
									Emph   struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"title"`
								Lb     []string `xml:"lb"`
								I      string   `xml:"i"`
								Name   []string `xml:"name"`
								Extref struct {
									Text    string `xml:",chardata"`
									Href    string `xml:"href,attr"`
									Archref string `xml:"archref"`
								} `xml:"extref"`
								Persname []string `xml:"persname"`
							} `xml:"unittitle"`
							Unitid []struct {
								Text       string `xml:",chardata"`
								Audience   string `xml:"audience,attr"`
								Identifier string `xml:"identifier,attr"`
								Type       string `xml:"type,attr"`
							} `xml:"unitid"`
							Physdesc []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Audience  string `xml:"audience,attr"`
								Extent    []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
									Lb        string `xml:"lb"`
								} `xml:"extent"`
								Dimensions struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"dimensions"`
								Physfacet struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"physfacet"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"physdesc"`
							Unitdate []struct {
								Text      string `xml:",chardata"`
								Normal    string `xml:"normal,attr"`
								Type      string `xml:"type,attr"`
								Certainty string `xml:"certainty,attr"`
							} `xml:"unitdate"`
							Container []struct {
								Text      string `xml:",chardata"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Type      string `xml:"type,attr"`
								Parent    string `xml:"parent,attr"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"container"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Label    string `xml:"label,attr"`
								Audience string `xml:"audience,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Role   string `xml:"role,attr"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
									Role   string `xml:"role,attr"`
									Rules  string `xml:"rules,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
							Langmaterial struct {
								Text     string `xml:",chardata"`
								ID       string `xml:"id,attr"`
								Label    string `xml:"label,attr"`
								Language struct {
									Text     string `xml:",chardata"`
									Langcode string `xml:"langcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
							Physloc struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
							} `xml:"physloc"`
							Materialspec struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"materialspec"`
							Abstract struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
							} `xml:"abstract"`
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
								P    []struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Name  []string `xml:"name"`
									Lb    []string `xml:"lb"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"p"`
							} `xml:"daodesc"`
						} `xml:"dao"`
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
							Geogname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"geogname"`
							Persname []struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
							Title []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"title"`
							Corpname []struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
							Function struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"function"`
							Occupation struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"occupation"`
						} `xml:"controlaccess"`
						Odd []struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Audience string `xml:"audience,attr"`
							Head     string `xml:"head"`
							P        []struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Lb     []string `xml:"lb"`
								Extent string   `xml:"extent"`
								Extref struct {
									Text    string `xml:",chardata"`
									Actuate string `xml:"actuate,attr"`
									Href    string `xml:"href,attr"`
									Type    string `xml:"type,attr"`
									Show    string `xml:"show,attr"`
								} `xml:"extref"`
								Blockquote string `xml:"blockquote"`
								Title      string `xml:"title"`
								Num        struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"num"`
							} `xml:"p"`
						} `xml:"odd"`
						C []struct {
							Text       string `xml:",chardata"`
							ID         string `xml:"id,attr"`
							Level      string `xml:"level,attr"`
							Otherlevel string `xml:"otherlevel,attr"`
							Audience   string `xml:"audience,attr"`
							Did        struct {
								Text      string `xml:",chardata"`
								Unittitle struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
										Title  string `xml:"title,attr"`
									} `xml:"emph"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"title"`
									Lb []string `xml:"lb"`
								} `xml:"unittitle"`
								Unitid []struct {
									Text       string `xml:",chardata"`
									Audience   string `xml:"audience,attr"`
									Identifier string `xml:"identifier,attr"`
									Type       string `xml:"type,attr"`
								} `xml:"unitid"`
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
								Langmaterial struct {
									Text     string `xml:",chardata"`
									Language struct {
										Text     string `xml:",chardata"`
										Langcode string `xml:"langcode,attr"`
									} `xml:"language"`
								} `xml:"langmaterial"`
								Physdesc []struct {
									Text      string `xml:",chardata"`
									ID        string `xml:"id,attr"`
									Altrender string `xml:"altrender,attr"`
									Label     string `xml:"label,attr"`
									Extent    []struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
									} `xml:"extent"`
								} `xml:"physdesc"`
								Origination []struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Audience string `xml:"audience,attr"`
									Corpname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"corpname"`
									Persname struct {
										Text   string `xml:",chardata"`
										Source string `xml:"source,attr"`
										Rules  string `xml:"rules,attr"`
									} `xml:"persname"`
								} `xml:"origination"`
								Physloc struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"physloc"`
							} `xml:"did"`
							Odd []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
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
									Extref struct {
										Text    string `xml:",chardata"`
										Type    string `xml:"type,attr"`
										Actuate string `xml:"actuate,attr"`
										Show    string `xml:"show,attr"`
										Href    string `xml:"href,attr"`
									} `xml:"extref"`
								} `xml:"p"`
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
										Emph struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"emph"`
									} `xml:"unittitle"`
									Langmaterial struct {
										Text     string `xml:",chardata"`
										Language struct {
											Text     string `xml:",chardata"`
											Langcode string `xml:"langcode,attr"`
										} `xml:"language"`
									} `xml:"langmaterial"`
									Container []struct {
										Text      string `xml:",chardata"`
										ID        string `xml:"id,attr"`
										Type      string `xml:"type,attr"`
										Label     string `xml:"label,attr"`
										Altrender string `xml:"altrender,attr"`
										Parent    string `xml:"parent,attr"`
									} `xml:"container"`
									Unitdate struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
									Unitid []struct {
										Text       string `xml:",chardata"`
										Audience   string `xml:"audience,attr"`
										Identifier string `xml:"identifier,attr"`
										Type       string `xml:"type,attr"`
									} `xml:"unitid"`
								} `xml:"did"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    string `xml:"p"`
								} `xml:"odd"`
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
									Text string   `xml:",chardata"`
									ID   string   `xml:"id,attr"`
									Head string   `xml:"head"`
									P    []string `xml:"p"`
								} `xml:"separatedmaterial"`
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
								Persname struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"persname"`
							} `xml:"controlaccess"`
							Userestrict struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"userestrict"`
							Phystech struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"phystech"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    []struct {
									Text string `xml:",chardata"`
									I    string `xml:"i"`
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
							Daogrp struct {
								Text    string `xml:",chardata"`
								Title   string `xml:"title,attr"`
								Type    string `xml:"type,attr"`
								Daodesc struct {
									Text string `xml:",chardata"`
									P    string `xml:"p"`
								} `xml:"daodesc"`
								Daoloc []struct {
									Text  string `xml:",chardata"`
									Href  string `xml:"href,attr"`
									Title string `xml:"title,attr"`
									Type  string `xml:"type,attr"`
								} `xml:"daoloc"`
							} `xml:"daogrp"`
							Accessrestrict struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"accessrestrict"`
							Separatedmaterial struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"separatedmaterial"`
							Bioghist struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"bioghist"`
							Relatedmaterial struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    []struct {
									Text   string `xml:",chardata"`
									Extref struct {
										Text    string `xml:",chardata"`
										Actuate string `xml:"actuate,attr"`
										Href    string `xml:"href,attr"`
									} `xml:"extref"`
								} `xml:"p"`
							} `xml:"relatedmaterial"`
						} `xml:"c"`
						Accessrestrict []struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Audience string `xml:"audience,attr"`
							Head     string `xml:"head"`
							P        []struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
							Legalstatus struct {
								Text     string `xml:",chardata"`
								Audience string `xml:"audience,attr"`
							} `xml:"legalstatus"`
						} `xml:"accessrestrict"`
						Phystech []struct {
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
							Lb []string `xml:"lb"`
						} `xml:"phystech"`
						Scopecontent []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
									Extref struct {
										Text string `xml:",chardata"`
										Href string `xml:"href,attr"`
									} `xml:"extref"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"emph"`
								Lb    []string `xml:"lb"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
								Ref []struct {
									Text   string `xml:",chardata"`
									Target string `xml:"target,attr"`
								} `xml:"ref"`
								Extref []struct {
									Text    string `xml:",chardata"`
									Href    string `xml:"href,attr"`
									Actuate string `xml:"actuate,attr"`
								} `xml:"extref"`
							} `xml:"p"`
							List struct {
								Text       string `xml:",chardata"`
								Numeration string `xml:"numeration,attr"`
								Type       string `xml:"type,attr"`
								Head       string `xml:"head"`
								Item       []struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"item"`
							} `xml:"list"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"scopecontent"`
						Separatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string   `xml:",chardata"`
								Lb   []string `xml:"lb"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Title struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
								Genreform struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
								Archref struct {
									Text    string `xml:",chardata"`
									Physloc string `xml:"physloc"`
								} `xml:"archref"`
							} `xml:"p"`
						} `xml:"separatedmaterial"`
						Processinfo struct {
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
						} `xml:"processinfo"`
						Altformavail struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
									Extref struct {
										Text    string `xml:",chardata"`
										Actuate string `xml:"actuate,attr"`
										Href    string `xml:"href,attr"`
									} `xml:"extref"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"altformavail"`
						Relatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string `xml:",chardata"`
								Ref  struct {
									Text   string `xml:",chardata"`
									Target string `xml:"target,attr"`
								} `xml:"ref"`
								Extref struct {
									Text string `xml:",chardata"`
									Href string `xml:"href,attr"`
								} `xml:"extref"`
							} `xml:"p"`
						} `xml:"relatedmaterial"`
						Bioghist struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text  string `xml:",chardata"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"p"`
						} `xml:"bioghist"`
						Acqinfo struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"acqinfo"`
						Userestrict struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"userestrict"`
						Arrangement struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"arrangement"`
						Otherfindaid struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"otherfindaid"`
						Fileplan struct {
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
						} `xml:"fileplan"`
						Originalsloc struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"originalsloc"`
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
						Custodhist struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"custodhist"`
					} `xml:"c"`
					Dao []struct {
						Text     string `xml:",chardata"`
						Actuate  string `xml:"actuate,attr"`
						Href     string `xml:"href,attr"`
						Role     string `xml:"role,attr"`
						Show     string `xml:"show,attr"`
						Title    string `xml:"title,attr"`
						Type     string `xml:"type,attr"`
						Audience string `xml:"audience,attr"`
						Daodesc  struct {
							Text string `xml:",chardata"`
							P    struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Name string `xml:"name"`
								Lb   string `xml:"lb"`
							} `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Controlaccess struct {
						Text    string `xml:",chardata"`
						Subject []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"subject"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Geogname []struct {
							Text           string `xml:",chardata"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"geogname"`
						Occupation []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"occupation"`
						Persname []struct {
							Text           string `xml:",chardata"`
							Authfilenumber string `xml:"authfilenumber,attr"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
						} `xml:"persname"`
						Corpname []struct {
							Text           string `xml:",chardata"`
							Authfilenumber string `xml:"authfilenumber,attr"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
						} `xml:"corpname"`
						Title []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"title"`
						Famname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
							Role   string `xml:"role,attr"`
						} `xml:"famname"`
					} `xml:"controlaccess"`
					Scopecontent []struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Audience string `xml:"audience,attr"`
						Head     struct {
							Text string `xml:",chardata"`
							Lb   string `xml:"lb"`
						} `xml:"head"`
						P []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
							} `xml:"emph"`
							Lb    []string `xml:"lb"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"title"`
							I      []string `xml:"i"`
							Extref []struct {
								Text    string `xml:",chardata"`
								Href    string `xml:"href,attr"`
								Actuate string `xml:"actuate,attr"`
							} `xml:"extref"`
							Ref []struct {
								Text   string `xml:",chardata"`
								Target string `xml:"target,attr"`
							} `xml:"ref"`
							Chronlist struct {
								Text      string `xml:",chardata"`
								Chronitem []struct {
									Text string `xml:",chardata"`
									Date struct {
										Text     string `xml:",chardata"`
										Calendar string `xml:"calendar,attr"`
										Era      string `xml:"era,attr"`
									} `xml:"date"`
									Event string `xml:"event"`
								} `xml:"chronitem"`
							} `xml:"chronlist"`
							Blockquote struct {
								Text string `xml:",chardata"`
								P    []struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Lb []string `xml:"lb"`
								} `xml:"p"`
							} `xml:"blockquote"`
							Archref struct {
								Text   string `xml:",chardata"`
								Type   string `xml:"type,attr"`
								Extref struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
									Href string `xml:"href,attr"`
								} `xml:"extref"`
							} `xml:"archref"`
						} `xml:"p"`
						Lb   []string `xml:"lb"`
						List struct {
							Text       string `xml:",chardata"`
							Type       string `xml:"type,attr"`
							Numeration string `xml:"numeration,attr"`
							Item       []struct {
								Text string `xml:",chardata"`
								Ref  struct {
									Text   string `xml:",chardata"`
									Target string `xml:"target,attr"`
								} `xml:"ref"`
							} `xml:"item"`
							Head string `xml:"head"`
						} `xml:"list"`
					} `xml:"scopecontent"`
					Odd []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Type   string `xml:"type,attr"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"title"`
							Lb     []string `xml:"lb"`
							Extref struct {
								Text    string `xml:",chardata"`
								Type    string `xml:"type,attr"`
								Actuate string `xml:"actuate,attr"`
								Show    string `xml:"show,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"odd"`
					Accessrestrict []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"head"`
						P []struct {
							Text   string `xml:",chardata"`
							Extref []struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
								Show    string `xml:"show,attr"`
							} `xml:"extref"`
							A struct {
								Text   string `xml:",chardata"`
								Href   string `xml:"href,attr"`
								Target string `xml:"target,attr"`
							} `xml:"a"`
						} `xml:"p"`
						Legalstatus struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"legalstatus"`
					} `xml:"accessrestrict"`
					Processinfo struct {
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
					} `xml:"processinfo"`
					Separatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string   `xml:",chardata"`
							Lb   []string `xml:"lb"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"separatedmaterial"`
					Userestrict struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"userestrict"`
					Relatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text   string `xml:",chardata"`
							Extref []struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Show    string `xml:"show,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
							Ref struct {
								Text   string `xml:",chardata"`
								Target string `xml:"target,attr"`
							} `xml:"ref"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"relatedmaterial"`
					Arrangement struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb    []string `xml:"lb"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"arrangement"`
					Custodhist struct {
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
					} `xml:"custodhist"`
					Phystech []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head struct {
							Text string `xml:",chardata"`
							Lb   string `xml:"lb"`
						} `xml:"head"`
						P []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
						Lb []string `xml:"lb"`
					} `xml:"phystech"`
					Altformavail []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text     string `xml:",chardata"`
								Linktype string `xml:"linktype,attr"`
								Actuate  string `xml:"actuate,attr"`
								Show     string `xml:"show,attr"`
								Href     string `xml:"href,attr"`
							} `xml:"extref"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Extref struct {
									Text    string `xml:",chardata"`
									Actuate string `xml:"actuate,attr"`
									Href    string `xml:"href,attr"`
								} `xml:"extref"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"altformavail"`
					Originalsloc struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"originalsloc"`
					Bioghist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"bioghist"`
					Otherfindaid struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Audience string `xml:"audience,attr"`
						Head     string `xml:"head"`
						P        []struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"otherfindaid"`
					Bibliography struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Extref string `xml:"extref"`
						} `xml:"p"`
					} `xml:"bibliography"`
					Acqinfo struct {
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
					} `xml:"acqinfo"`
					Appraisal struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"appraisal"`
					Fileplan struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"fileplan"`
					Accruals struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accruals"`
					Index struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						Head       string `xml:"head"`
						Indexentry []struct {
							Text     string `xml:",chardata"`
							Name     string `xml:"name"`
							Subject  string `xml:"subject"`
							Corpname string `xml:"corpname"`
						} `xml:"indexentry"`
					} `xml:"index"`
					Prefercite struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"prefercite"`
				} `xml:"c"`
				Scopecontent []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text      string `xml:",chardata"`
							Render    string `xml:"render,attr"`
							Renderr   string `xml:"renderr,attr"`
							AttrTitle string `xml:"title,attr"`
							Extref    struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"title"`
						} `xml:"emph"`
						Lb       []string `xml:"lb"`
						Corpname string   `xml:"corpname"`
						Date     []struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Subject    []string `xml:"subject"`
						Genreform  []string `xml:"genreform"`
						Occupation string   `xml:"occupation"`
						Title      []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
							Emph   struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"title"`
						I      string `xml:"i"`
						Extref []struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
							Emph    struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"extref"`
						Ref []struct {
							Text   string `xml:",chardata"`
							Target string `xml:"target,attr"`
						} `xml:"ref"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text    string `xml:",chardata"`
								Href    string `xml:"href,attr"`
								Type    string `xml:"type,attr"`
								Actuate string `xml:"actuate,attr"`
								Show    string `xml:"show,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
						Blockquote struct {
							Text string `xml:",chardata"`
							P    struct {
								Text   string `xml:",chardata"`
								Extref struct {
									Text    string `xml:",chardata"`
									Actuate string `xml:"actuate,attr"`
									Href    string `xml:"href,attr"`
									Title   []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"extref"`
							} `xml:"p"`
						} `xml:"blockquote"`
						List struct {
							Text string `xml:",chardata"`
							Item []struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"item"`
						} `xml:"list"`
					} `xml:"p"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					List struct {
						Text       string `xml:",chardata"`
						Numeration string `xml:"numeration,attr"`
						Type       string `xml:"type,attr"`
						Head       string `xml:"head"`
						Item       []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"item"`
						Defitem struct {
							Text  string `xml:",chardata"`
							Label string `xml:"label"`
							Item  string `xml:"item"`
						} `xml:"defitem"`
					} `xml:"list"`
					Lb     []string `xml:"lb"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"scopecontent"`
				Accessrestrict struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
							Show    string `xml:"show,attr"`
						} `xml:"extref"`
					} `xml:"p"`
				} `xml:"accessrestrict"`
				Arrangement struct {
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
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Ref struct {
							Text   string `xml:",chardata"`
							Target string `xml:"target,attr"`
						} `xml:"ref"`
						P          string `xml:"p"`
						Blockquote struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"blockquote"`
					} `xml:"p"`
					List struct {
						Text       string `xml:",chardata"`
						Type       string `xml:"type,attr"`
						Numeration string `xml:"numeration,attr"`
						Defitem    []struct {
							Text  string `xml:",chardata"`
							Label string `xml:"label"`
							Item  string `xml:"item"`
						} `xml:"defitem"`
						Head string   `xml:"head"`
						Item []string `xml:"item"`
					} `xml:"list"`
				} `xml:"arrangement"`
				Bioghist []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Audience string `xml:"audience,attr"`
					Head     string `xml:"head"`
					P        []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						I     []string `xml:"i"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Lb []string `xml:"lb"`
					} `xml:"p"`
				} `xml:"bioghist"`
				Odd []struct {
					Text     string `xml:",chardata"`
					ID       string `xml:"id,attr"`
					Audience string `xml:"audience,attr"`
					Head     string `xml:"head"`
					P        []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb     []string `xml:"lb"`
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
						Ref struct {
							Text    string `xml:",chardata"`
							Href    string `xml:"href,attr"`
							Show    string `xml:"show,attr"`
							Actuate string `xml:"actuate,attr"`
						} `xml:"ref"`
					} `xml:"p"`
				} `xml:"odd"`
				Processinfo struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"processinfo"`
				Userestrict struct {
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
				} `xml:"userestrict"`
				Phystech []struct {
					Text     string   `xml:",chardata"`
					ID       string   `xml:"id,attr"`
					Audience string   `xml:"audience,attr"`
					Head     string   `xml:"head"`
					P        []string `xml:"p"`
				} `xml:"phystech"`
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
				Otherfindaid struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"otherfindaid"`
				Custodhist struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"custodhist"`
				Acqinfo struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"acqinfo"`
				Altformavail struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"emph"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"p"`
				} `xml:"altformavail"`
				Relatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text   string `xml:",chardata"`
						Extref struct {
							Text    string `xml:",chardata"`
							Href    string `xml:"href,attr"`
							Actuate string `xml:"actuate,attr"`
						} `xml:"extref"`
						Emph string `xml:"emph"`
					} `xml:"p"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"relatedmaterial"`
				Separatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"p"`
				} `xml:"separatedmaterial"`
				Prefercite struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"prefercite"`
				Daogrp struct {
					Text    string `xml:",chardata"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Daodesc struct {
						Text string `xml:",chardata"`
						P    string `xml:"p"`
					} `xml:"daodesc"`
					Daoloc []struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Role  string `xml:"role,attr"`
						Title string `xml:"title,attr"`
						Type  string `xml:"type,attr"`
					} `xml:"daoloc"`
				} `xml:"daogrp"`
				Originalsloc struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"originalsloc"`
			} `xml:"c"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"dsc"`
		Bioghist []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				List []struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Archref struct {
							Text   string `xml:",chardata"`
							Extref struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"archref"`
						Bibref struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"bibref"`
					} `xml:"item"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"list"`
				Corpname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"corpname"`
				Subject  []string `xml:"subject"`
				Persname []struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"persname"`
				Geogname []string `xml:"geogname"`
				Emph     []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Extref    struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Date []struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				Occupation []string `xml:"occupation"`
				Name       []string `xml:"name"`
				Extref     []struct {
					Text     string `xml:",chardata"`
					Href     string `xml:"href,attr"`
					Actuate  string `xml:"actuate,attr"`
					Linktype string `xml:"linktype,attr"`
					Show     string `xml:"show,attr"`
					Title    string `xml:"title,attr"`
					Type     string `xml:"type,attr"`
					Emph     struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				Lb         []string `xml:"lb"`
				Blockquote struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"blockquote"`
				Abbr  string `xml:"abbr"`
				Table struct {
					Text   string `xml:",chardata"`
					Tgroup struct {
						Text    string `xml:",chardata"`
						Cols    string `xml:"cols,attr"`
						Colspec []struct {
							Text     string `xml:",chardata"`
							Colwidth string `xml:"colwidth,attr"`
						} `xml:"colspec"`
						Tbody struct {
							Text string `xml:",chardata"`
							Row  []struct {
								Text  string `xml:",chardata"`
								Entry []struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"entry"`
							} `xml:"row"`
						} `xml:"tbody"`
					} `xml:"tgroup"`
				} `xml:"table"`
				Chronlist struct {
					Text      string `xml:",chardata"`
					Chronitem []struct {
						Text string `xml:",chardata"`
						Date struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Event struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"event"`
						Eventgrp struct {
							Text  string `xml:",chardata"`
							Event []struct {
								Text  string `xml:",chardata"`
								Title struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"event"`
						} `xml:"eventgrp"`
					} `xml:"chronitem"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"chronlist"`
				Archref struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Type string `xml:"type,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Bibref struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"bibref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Item       []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
					} `xml:"extref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"item"`
				Head    string `xml:"head"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
			Chronlist []struct {
				Text     string `xml:",chardata"`
				Audience string `xml:"audience,attr"`
				Head     struct {
					Text  string `xml:",chardata"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"head"`
				Chronitem []struct {
					Text     string `xml:",chardata"`
					Date     string `xml:"date"`
					Eventgrp struct {
						Text  string `xml:",chardata"`
						Event []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"title"`
							Lb     []string `xml:"lb"`
							Extref struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
							} `xml:"extref"`
						} `xml:"event"`
					} `xml:"eventgrp"`
					Event struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Extref struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
						} `xml:"extref"`
						Lb   string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"event"`
				} `xml:"chronitem"`
			} `xml:"chronlist"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Extref []struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"extref"`
			Lb []string `xml:"lb"`
		} `xml:"bioghist"`
		Arrangement []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				List struct {
					Text       string `xml:",chardata"`
					Numeration string `xml:"numeration,attr"`
					Type       string `xml:"type,attr"`
					Item       []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"item"`
				} `xml:"list"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Title  string `xml:"title,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Lb []string `xml:"lb"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Audience   string `xml:"audience,attr"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Title  struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
					Ref struct {
						Text   string `xml:",chardata"`
						Target string `xml:"target,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"ref"`
					Lb []string `xml:"lb"`
				} `xml:"item"`
				Head    string `xml:"head"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
		} `xml:"arrangement"`
		Accessrestrict []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Type    string `xml:"type,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
				} `xml:"extref"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
			Legalstatus string `xml:"legalstatus"`
		} `xml:"accessrestrict"`
		Custodhist []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Lb   []string `xml:"lb"`
				Date struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"custodhist"`
		Userestrict []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text    string   `xml:",chardata"`
				Lb      []string `xml:"lb"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Corpname string `xml:"corpname"`
			} `xml:"p"`
		} `xml:"userestrict"`
		Acqinfo []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb     []string `xml:"lb"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"acqinfo"`
		Prefercite []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text    string   `xml:",chardata"`
				Lb      []string `xml:"lb"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string `xml:",chardata"`
						Lb   string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Odd []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				List struct {
					Text string   `xml:",chardata"`
					Item []string `xml:"item"`
				} `xml:"list"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text     string   `xml:",chardata"`
					Persname []string `xml:"persname"`
					Corpname string   `xml:"corpname"`
					Emph     struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Title  struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Render string `xml:"render,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"odd"`
		Processinfo []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     string `xml:"head"`
			P        []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb     []string `xml:"lb"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
				} `xml:"extref"`
				I     []string `xml:"i"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"processinfo"`
		Phystech []struct {
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
		} `xml:"phystech"`
		Separatedmaterial []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Audience       string `xml:"audience,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Archref struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
						Type string `xml:"type,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				List []struct {
					Text string `xml:",chardata"`
					Head string `xml:"head"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Lb []string `xml:"lb"`
					} `xml:"item"`
				} `xml:"list"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
				Head string `xml:"head"`
			} `xml:"list"`
		} `xml:"separatedmaterial"`
		Relatedmaterial []struct {
			Text           string `xml:",chardata"`
			ID             string `xml:"id,attr"`
			Encodinganalog string `xml:"encodinganalog,attr"`
			Head           string `xml:"head"`
			P              []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Href    string `xml:"href,attr"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Xlink   string `xml:"xlink,attr"`
					Target  string `xml:"target,attr"`
					Archref struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"archref"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Title  string `xml:"title,attr"`
					} `xml:"emph"`
				} `xml:"extref"`
				Archref []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Type    string `xml:"type,attr"`
						Href    string `xml:"href,attr"`
						Show    string `xml:"show,attr"`
						Actuate string `xml:"actuate,attr"`
						Emph    struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"extref"`
				} `xml:"archref"`
				Emph []struct {
					Text      string `xml:",chardata"`
					Render    string `xml:"render,attr"`
					AttrTitle string `xml:"title,attr"`
					Extref    struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"title"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				List struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
					} `xml:"item"`
				} `xml:"list"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Audience   string `xml:"audience,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text    string `xml:",chardata"`
					Archref struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Extref struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"archref"`
					List struct {
						Text string   `xml:",chardata"`
						Item []string `xml:"item"`
					} `xml:"list"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Extref struct {
						Text    string `xml:",chardata"`
						Href    string `xml:"href,attr"`
						Actuate string `xml:"actuate,attr"`
						Show    string `xml:"show,attr"`
					} `xml:"extref"`
					Title []struct {
						Text   string `xml:",chardata"`
						Type   string `xml:"type,attr"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"item"`
			} `xml:"list"`
			Emph []struct {
				Text   string `xml:",chardata"`
				Render string `xml:"render,attr"`
			} `xml:"emph"`
			Lb     []string `xml:"lb"`
			Extref []struct {
				Text    string `xml:",chardata"`
				Href    string `xml:"href,attr"`
				Actuate string `xml:"actuate,attr"`
				Emph    struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"extref"`
		} `xml:"relatedmaterial"`
		Appraisal []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Rend   string `xml:"rend,attr"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"appraisal"`
		Originalsloc struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"originalsloc"`
		Otherfindaid struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Archref struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
			} `xml:"p"`
		} `xml:"otherfindaid"`
		Accruals struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accruals"`
		Altformavail struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text   string `xml:",chardata"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Extref struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"altformavail"`
		Bibliography struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Audience string `xml:"audience,attr"`
			Head     string `xml:"head"`
			Bibref   []struct {
				Text  string `xml:",chardata"`
				Title []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"bibref"`
			P []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Bibref []struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
				} `xml:"bibref"`
			} `xml:"p"`
		} `xml:"bibliography"`
		Index struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Head       string `xml:"head"`
			P          string `xml:"p"`
			Indexentry []struct {
				Text  string `xml:",chardata"`
				Title string `xml:"title"`
				Ref   string `xml:"ref"`
			} `xml:"indexentry"`
		} `xml:"index"`
		Dao struct {
			Text    string `xml:",chardata"`
			Actuate string `xml:"actuate,attr"`
			Show    string `xml:"show,attr"`
			Title   string `xml:"title,attr"`
			Role    string `xml:"role,attr"`
			Href    string `xml:"href,attr"`
			Daodesc struct {
				Text string `xml:",chardata"`
				P    string `xml:"p"`
			} `xml:"daodesc"`
		} `xml:"dao"`
	} `xml:"archdesc"`
} 

