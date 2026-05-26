package carapace

type (
	batch        []Action
	invokedBatch []InvokedAction
)

// Batch creates a batch of Actions that can be invoked in parallel.
func Batch(actions ...Action) batch {
	_ = "STUB: not implemented"
	return *

	// Invoke invokes contained Actions of the batch using goroutines.
	new(batch)
}

func (b batch) Invoke(c Context) invokedBatch { _ = "STUB: not implemented"; return *new(invokedBatch) }

// ToA converts the batch to an implicitly merged action which is a shortcut for:
//
//	ActionCallback(func(c Context) Action {
//		return batch.Invoke(c).Merge().ToA()
//	})
func (b batch) ToA() Action { _ = "STUB: not implemented"; return *new(Action) }

// Merge merges Actions of a batch.
func (b invokedBatch) Merge() InvokedAction { _ = "STUB: not implemented"; return *new(InvokedAction) }

// Parallelize parallelizes the function calls (https://stackoverflow.com/a/44402936)
func parallelize(functions ...func()) { _ = "STUB: not implemented"; return }
