package chordshapes

func GetVoicings(root Note, quality ChordQuality) []Voicing {
	var voicings []Voicing
	for _, template := range shapeTemplates {
		if template.Quality != quality {
			continue
		}

		if template.BaseRoot == root && !template.Movable {
			voicings = append(voicings, Voicing{
				Name:    root.String() + string(quality),
				Root:    root,
				Quality: quality,
				Frets:   template.Frets,
				Fingers: template.Fingers,
				Source:  template.Name,
				Tags:    append([]string(nil), template.Tags...),
			})
			continue
		}

		voicing, err := TransposeShape(template, root)
		if err != nil {
			continue
		}
		voicings = append(voicings, voicing)
	}

	return voicings
}
