package main

import "fmt"

func main() {

	b := []byte{57, 0, 56, 0, 55, 0, 54, 0, 53, 0, 52, 0, 51, 0, 50, 0, 49, 0, 48}

	for i := 0; i < 3; i++ {
		if i == 0 {
			b[1] = 0
		} else if i == 1 {
			b[1] = 45
		} else if i == 2 {
			b[1] = 43
		}

		for i := 0; i < 3; i++ {
			if i == 0 {
				b[3] = 0
			} else if i == 1 {
				b[3] = 45
			} else if i == 2 {
				b[3] = 43
			}

			for i := 0; i < 3; i++ {
				if i == 0 {
					b[5] = 0
				} else if i == 1 {
					b[5] = 45
				} else if i == 2 {
					b[5] = 43
				}

				for i := 0; i < 3; i++ {
					if i == 0 {
						b[7] = 0
					} else if i == 1 {
						b[7] = 45
					} else if i == 2 {
						b[7] = 43
					}

					for i := 0; i < 3; i++ {
						if i == 0 {
							b[9] = 0
						} else if i == 1 {
							b[9] = 45
						} else if i == 2 {
							b[9] = 43
						}

						for i := 0; i < 3; i++ {
							if i == 0 {
								b[11] = 0
							} else if i == 1 {
								b[11] = 45
							} else if i == 2 {
								b[11] = 43
							}

							for i := 0; i < 3; i++ {
								if i == 0 {
									b[13] = 0
								} else if i == 1 {
									b[13] = 45
								} else if i == 2 {
									b[13] = 43
								}

								for i := 0; i < 3; i++ {
									if i == 0 {
										b[15] = 0
									} else if i == 1 {
										b[15] = 45
									} else if i == 2 {
										b[15] = 43
									}

									for i := 0; i < 3; i++ {
										if i == 0 {
											b[17] = 0
										} else if i == 1 {
											b[17] = 45
										} else if i == 2 {
											b[17] = 43
										}

										a := convert(b)
										if a == 200 {
											fmt.Println(string(b) + " = 200")
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func convert(b []byte) int {

	sum := 0 //итоговая сумма
	a := 0   //промежуточная сумма
	z := 43  //индикатор знака(для начала - "+")

	//перебираем строку
	for i := 0; i <= len(b)-1; i++ {

		if b[i] != 43 && b[i] != 45 && b[i] != 0 {
			a = a*10 + (int(b[i]) - 48) //так проще
		}

		if b[i] == 43 || b[i] == 45 || i == len(b)-1 {
			if z == 43 {
				sum = sum + a
			} else {
				sum = sum - a
			}
			a = 0
			z = int(b[i])
		}

	}
	return (sum)
}
