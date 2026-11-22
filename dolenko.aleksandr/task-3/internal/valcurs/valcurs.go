func Load(path string) (ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ValCurs{}, fmt.Errorf("read XML file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var currencies ValCurs

	err = decoder.Decode(&currencies)
	if err != nil {
		return ValCurs{}, fmt.Errorf("decode XML: %w", err)
	}

	if len(currencies.Valutes) == 0 {
		return ValCurs{}, fmt.Errorf("XML contains no currencies")
	}

	for i, currency := range currencies.Valutes {
		if currency.NumCode <= 0 {
			return ValCurs{}, fmt.Errorf("invalid NumCode at currency %d", i+1)
		}

		if strings.TrimSpace(currency.CharCode) == "" {
			return ValCurs{}, fmt.Errorf("empty CharCode at currency %d", i+1)
		}

		if currency.Value <= 0 {
			return ValCurs{}, fmt.Errorf("invalid Value at currency %d", i+1)
		}
	}

	return currencies, nil
}
