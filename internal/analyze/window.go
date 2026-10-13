package analyze

// RepetitionWindow reports how often a token repeats inside a sliding window.
// It is a sliding counter: callers advance the window instead of rescanning the
// whole file for every element, which keeps the repetition pass linear in the
// number of lines rather than quadratic.
type RepetitionWindow struct {
	size   int
	counts map[string]int
	order  []string
	total  int
}

// NewRepetitionWindow returns a window of the given size. A non-positive size
// is corrected to 1 rather than producing a window that can never match.
func NewRepetitionWindow(size int) *RepetitionWindow {
	if size < 1 {
		size = 1
	}
	return &RepetitionWindow{size: size, counts: make(map[string]int, size)}
}

// Count returns how many times token appears in the current window.
func (w *RepetitionWindow) Count(token string) int {
	return w.counts[token]
}

// Push adds a token, evicting the oldest one once the window is full.
func (w *RepetitionWindow) Push(token string) {
	w.counts[token]++
	w.order = append(w.order, token)
	w.total++
	if len(w.order) > w.size {
		oldest := w.order[0]
		w.order = w.order[1:]
		w.counts[oldest]--
		if w.counts[oldest] == 0 {
			delete(w.counts, oldest)
		}
	}
}

// Full reports whether the window has reached its size.
func (w *RepetitionWindow) Full() bool { return len(w.order) == w.size }
