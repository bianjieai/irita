package errors

import "fmt"

const RootCodeSpace = "relayer"

var (
	ErrInternal                     = Register(RootCodeSpace, 1, "internal")
	ErrChainConn                    = Register(RootCodeSpace, 2, "connection chain failed")
	ErrUnknown                      = Register(RootCodeSpace, 3, "failed to unknown ")
	ErrGetRandomWordsRequestedEvent = Register(RootCodeSpace, 4, "failed to get random words requested event")
	ErrGetLatestHeight              = Register(RootCodeSpace, 5, "failed to get latest height")
	ErrFulfillRandomWords           = Register(RootCodeSpace, 6, "failed fulfill random words")
	ErrNotProduced                  = Register(RootCodeSpace, 7, "failed to not produced")
	ErrStoreBlockHash               = Register(RootCodeSpace, 8, "failed to store block hash")
)

var usedCodes = map[string]*Error{}

func getUsed(codespace string, code uint32) *Error {
	return usedCodes[errorID(codespace, code)]
}

func setUsed(err *Error) {
	usedCodes[errorID(err.codeSpace, err.code)] = err
}

func errorID(codespace string, code uint32) string {
	return fmt.Sprintf("%s:%d", codespace, code)
}

type IError interface {
	error
	Code() uint32
	CodeSpace() string
}

type Error struct {
	codeSpace string
	code      uint32
	desc      string
}

func New(codeSpace string, code uint32, desc string) *Error {
	return &Error{codeSpace: codeSpace, code: code, desc: desc}
}

func (e Error) Error() string {
	return e.desc
}

func (e Error) Code() uint32 {
	return e.code
}

func (e Error) CodeSpace() string {
	return e.codeSpace
}

func Register(codespace string, code uint32, description string) *Error {
	if e := getUsed(codespace, code); e != nil {
		panic(fmt.Sprintf("error with code %d is already registered: %q", code, e.desc))
	}

	err := New(codespace, code, description)
	setUsed(err)

	return err
}
