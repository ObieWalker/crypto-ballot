package errors

type PollingUnitNotExistError struct{}
type VoterAlreadyVotedError struct{}

func (*PollingUnitNotExistError) Error() string {
  return "This Polling Unit doesn't exists."
}

func (*VoterAlreadyVotedError) Error() string {
  return "A voter with this ID has already voted at this polling unit."
}