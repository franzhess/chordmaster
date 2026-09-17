package chordshapes

import "fmt"

func TransposeShape(shape ShapeTemplate, targetRoot Note) (Voicing, error) {
	distance := SemitoneDistance(shape.BaseRoot, targetRoot)
	frets := shape.Frets

	if shape.Movable {
		for i, fret := range frets {
			if fret == -1 {
				continue
			}
			if fret < -1 {
				return Voicing{}, fmt.Errorf("invalid negative fret %d", fret)
			}

			frets[i] = fret + distance
		}
	} else if shape.BaseRoot != targetRoot {
		return Voicing{}, fmt.Errorf("shape %q is not movable", shape.Name)
	}

	if err := validateFrets(frets); err != nil {
		return Voicing{}, err
	}

	return Voicing{
		Name:    targetRoot.String() + string(shape.Quality),
		Root:    targetRoot,
		Quality: shape.Quality,
		Frets:   frets,
		Fingers: shape.Fingers,
		Source:  shape.Name,
		Tags:    append([]string(nil), shape.Tags...),
	}, nil
}

func validateFrets(frets [6]int) error {
	for _, fret := range frets {
		if fret < -1 {
			return fmt.Errorf("invalid fret %d", fret)
		}
		if fret > 24 {
			return fmt.Errorf("fret %d is above 24", fret)
		}
	}

	return nil
}
