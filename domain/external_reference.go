package domain

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ExternalReferenceNamespace struct {
	value string
}

var ExternalReferenceNamespacePatient = ExternalReferenceNamespace{"patient"}

type ExternalReference struct {
	namespace ExternalReferenceNamespace
	id        string
}

func NewExternalReference(ns ExternalReferenceNamespace, id string) (*ExternalReference, error) {
	if ns.value == "" {
		return nil, errors.New("external reference namespace is empty")
	}
	if id == "" {
		return nil, errors.New("external reference id is empty")
	}
	return &ExternalReference{
		namespace: ns,
		id:        id,
	}, nil
}

func (r ExternalReference) String() string {
	return r.namespace.value + ":" + r.id
}

// db interface methods
func (r *ExternalReference) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal ExternalReference value:", value))
	}

	splits := strings.Split(s, ":")
	if len(splits) != 2 {
		return errors.New("unexpected format of ExternalReference")
	}
	namespace := splits[0]
	id := splits[1]

	switch namespace {
	case ExternalReferenceNamespacePatient.value:
		r.namespace = ExternalReferenceNamespacePatient
		r.id = id
	default:
		return errors.New("unknown ExternalReference namespace")
	}

	return nil
}

func (r ExternalReference) Value() (driver.Value, error) {
	if r.namespace.value == "" {
		return "", errors.New("no namespace set")
	}
	if r.id == "" {
		return "", errors.New("no reference id set")
	}

	return r.String(), nil
}

// compiler checks
var _ sql.Scanner = (*ExternalReference)(nil)
var _ driver.Valuer = (*ExternalReference)(nil)
var _ fmt.Stringer = (*ExternalReference)(nil)
