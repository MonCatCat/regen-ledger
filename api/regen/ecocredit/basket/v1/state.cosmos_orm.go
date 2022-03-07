// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package basketv1

import (
	context "context"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type BasketStore interface {
	Insert(ctx context.Context, basket *Basket) error
	InsertReturningID(ctx context.Context, basket *Basket) (uint64, error)
	Update(ctx context.Context, basket *Basket) error
	Save(ctx context.Context, basket *Basket) error
	Delete(ctx context.Context, basket *Basket) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Basket, error)
	HasByBasketDenom(ctx context.Context, basket_denom string) (found bool, err error)
	// GetByBasketDenom returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByBasketDenom(ctx context.Context, basket_denom string) (*Basket, error)
	HasByName(ctx context.Context, name string) (found bool, err error)
	// GetByName returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByName(ctx context.Context, name string) (*Basket, error)
	List(ctx context.Context, prefixKey BasketIndexKey, opts ...ormlist.Option) (BasketIterator, error)
	ListRange(ctx context.Context, from, to BasketIndexKey, opts ...ormlist.Option) (BasketIterator, error)
	DeleteBy(ctx context.Context, prefixKey BasketIndexKey) error
	DeleteRange(ctx context.Context, from, to BasketIndexKey) error

	doNotImplement()
}

type BasketIterator struct {
	ormtable.Iterator
}

func (i BasketIterator) Value() (*Basket, error) {
	var basket Basket
	err := i.UnmarshalMessage(&basket)
	return &basket, err
}

type BasketIndexKey interface {
	id() uint32
	values() []interface{}
	basketIndexKey()
}

// primary key starting index..
type BasketPrimaryKey = BasketIdIndexKey

type BasketIdIndexKey struct {
	vs []interface{}
}

func (x BasketIdIndexKey) id() uint32            { return 0 }
func (x BasketIdIndexKey) values() []interface{} { return x.vs }
func (x BasketIdIndexKey) basketIndexKey()       {}

func (this BasketIdIndexKey) WithId(id uint64) BasketIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type BasketBasketDenomIndexKey struct {
	vs []interface{}
}

func (x BasketBasketDenomIndexKey) id() uint32            { return 1 }
func (x BasketBasketDenomIndexKey) values() []interface{} { return x.vs }
func (x BasketBasketDenomIndexKey) basketIndexKey()       {}

func (this BasketBasketDenomIndexKey) WithBasketDenom(basket_denom string) BasketBasketDenomIndexKey {
	this.vs = []interface{}{basket_denom}
	return this
}

type BasketNameIndexKey struct {
	vs []interface{}
}

func (x BasketNameIndexKey) id() uint32            { return 2 }
func (x BasketNameIndexKey) values() []interface{} { return x.vs }
func (x BasketNameIndexKey) basketIndexKey()       {}

func (this BasketNameIndexKey) WithName(name string) BasketNameIndexKey {
	this.vs = []interface{}{name}
	return this
}

type basketStore struct {
	table ormtable.AutoIncrementTable
}

func (this basketStore) Insert(ctx context.Context, basket *Basket) error {
	return this.table.Insert(ctx, basket)
}

func (this basketStore) Update(ctx context.Context, basket *Basket) error {
	return this.table.Update(ctx, basket)
}

func (this basketStore) Save(ctx context.Context, basket *Basket) error {
	return this.table.Save(ctx, basket)
}

func (this basketStore) Delete(ctx context.Context, basket *Basket) error {
	return this.table.Delete(ctx, basket)
}

func (this basketStore) InsertReturningID(ctx context.Context, basket *Basket) (uint64, error) {
	return this.table.InsertReturningID(ctx, basket)
}

func (this basketStore) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this basketStore) Get(ctx context.Context, id uint64) (*Basket, error) {
	var basket Basket
	found, err := this.table.PrimaryKey().Get(ctx, &basket, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &basket, nil
}

func (this basketStore) HasByBasketDenom(ctx context.Context, basket_denom string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		basket_denom,
	)
}

func (this basketStore) GetByBasketDenom(ctx context.Context, basket_denom string) (*Basket, error) {
	var basket Basket
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &basket,
		basket_denom,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &basket, nil
}

func (this basketStore) HasByName(ctx context.Context, name string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		name,
	)
}

func (this basketStore) GetByName(ctx context.Context, name string) (*Basket, error) {
	var basket Basket
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &basket,
		name,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &basket, nil
}

func (this basketStore) List(ctx context.Context, prefixKey BasketIndexKey, opts ...ormlist.Option) (BasketIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BasketIterator{it}, err
}

func (this basketStore) ListRange(ctx context.Context, from, to BasketIndexKey, opts ...ormlist.Option) (BasketIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BasketIterator{it}, err
}

func (this basketStore) DeleteBy(ctx context.Context, prefixKey BasketIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this basketStore) DeleteRange(ctx context.Context, from, to BasketIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this basketStore) doNotImplement() {}

var _ BasketStore = basketStore{}

func NewBasketStore(db ormtable.Schema) (BasketStore, error) {
	table := db.GetTable(&Basket{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Basket{}).ProtoReflect().Descriptor().FullName()))
	}
	return basketStore{table.(ormtable.AutoIncrementTable)}, nil
}

type BasketClassStore interface {
	Insert(ctx context.Context, basketClass *BasketClass) error
	Update(ctx context.Context, basketClass *BasketClass) error
	Save(ctx context.Context, basketClass *BasketClass) error
	Delete(ctx context.Context, basketClass *BasketClass) error
	Has(ctx context.Context, basket_id uint64, class_id string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, basket_id uint64, class_id string) (*BasketClass, error)
	List(ctx context.Context, prefixKey BasketClassIndexKey, opts ...ormlist.Option) (BasketClassIterator, error)
	ListRange(ctx context.Context, from, to BasketClassIndexKey, opts ...ormlist.Option) (BasketClassIterator, error)
	DeleteBy(ctx context.Context, prefixKey BasketClassIndexKey) error
	DeleteRange(ctx context.Context, from, to BasketClassIndexKey) error

	doNotImplement()
}

type BasketClassIterator struct {
	ormtable.Iterator
}

func (i BasketClassIterator) Value() (*BasketClass, error) {
	var basketClass BasketClass
	err := i.UnmarshalMessage(&basketClass)
	return &basketClass, err
}

type BasketClassIndexKey interface {
	id() uint32
	values() []interface{}
	basketClassIndexKey()
}

// primary key starting index..
type BasketClassPrimaryKey = BasketClassBasketIdClassIdIndexKey

type BasketClassBasketIdClassIdIndexKey struct {
	vs []interface{}
}

func (x BasketClassBasketIdClassIdIndexKey) id() uint32            { return 0 }
func (x BasketClassBasketIdClassIdIndexKey) values() []interface{} { return x.vs }
func (x BasketClassBasketIdClassIdIndexKey) basketClassIndexKey()  {}

func (this BasketClassBasketIdClassIdIndexKey) WithBasketId(basket_id uint64) BasketClassBasketIdClassIdIndexKey {
	this.vs = []interface{}{basket_id}
	return this
}

func (this BasketClassBasketIdClassIdIndexKey) WithBasketIdClassId(basket_id uint64, class_id string) BasketClassBasketIdClassIdIndexKey {
	this.vs = []interface{}{basket_id, class_id}
	return this
}

type basketClassStore struct {
	table ormtable.Table
}

func (this basketClassStore) Insert(ctx context.Context, basketClass *BasketClass) error {
	return this.table.Insert(ctx, basketClass)
}

func (this basketClassStore) Update(ctx context.Context, basketClass *BasketClass) error {
	return this.table.Update(ctx, basketClass)
}

func (this basketClassStore) Save(ctx context.Context, basketClass *BasketClass) error {
	return this.table.Save(ctx, basketClass)
}

func (this basketClassStore) Delete(ctx context.Context, basketClass *BasketClass) error {
	return this.table.Delete(ctx, basketClass)
}

func (this basketClassStore) Has(ctx context.Context, basket_id uint64, class_id string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, basket_id, class_id)
}

func (this basketClassStore) Get(ctx context.Context, basket_id uint64, class_id string) (*BasketClass, error) {
	var basketClass BasketClass
	found, err := this.table.PrimaryKey().Get(ctx, &basketClass, basket_id, class_id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &basketClass, nil
}

func (this basketClassStore) List(ctx context.Context, prefixKey BasketClassIndexKey, opts ...ormlist.Option) (BasketClassIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BasketClassIterator{it}, err
}

func (this basketClassStore) ListRange(ctx context.Context, from, to BasketClassIndexKey, opts ...ormlist.Option) (BasketClassIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BasketClassIterator{it}, err
}

func (this basketClassStore) DeleteBy(ctx context.Context, prefixKey BasketClassIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this basketClassStore) DeleteRange(ctx context.Context, from, to BasketClassIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this basketClassStore) doNotImplement() {}

var _ BasketClassStore = basketClassStore{}

func NewBasketClassStore(db ormtable.Schema) (BasketClassStore, error) {
	table := db.GetTable(&BasketClass{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&BasketClass{}).ProtoReflect().Descriptor().FullName()))
	}
	return basketClassStore{table}, nil
}

type BasketBalanceStore interface {
	Insert(ctx context.Context, basketBalance *BasketBalance) error
	Update(ctx context.Context, basketBalance *BasketBalance) error
	Save(ctx context.Context, basketBalance *BasketBalance) error
	Delete(ctx context.Context, basketBalance *BasketBalance) error
	Has(ctx context.Context, basket_id uint64, batch_denom string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, basket_id uint64, batch_denom string) (*BasketBalance, error)
	List(ctx context.Context, prefixKey BasketBalanceIndexKey, opts ...ormlist.Option) (BasketBalanceIterator, error)
	ListRange(ctx context.Context, from, to BasketBalanceIndexKey, opts ...ormlist.Option) (BasketBalanceIterator, error)
	DeleteBy(ctx context.Context, prefixKey BasketBalanceIndexKey) error
	DeleteRange(ctx context.Context, from, to BasketBalanceIndexKey) error

	doNotImplement()
}

type BasketBalanceIterator struct {
	ormtable.Iterator
}

func (i BasketBalanceIterator) Value() (*BasketBalance, error) {
	var basketBalance BasketBalance
	err := i.UnmarshalMessage(&basketBalance)
	return &basketBalance, err
}

type BasketBalanceIndexKey interface {
	id() uint32
	values() []interface{}
	basketBalanceIndexKey()
}

// primary key starting index..
type BasketBalancePrimaryKey = BasketBalanceBasketIdBatchDenomIndexKey

type BasketBalanceBasketIdBatchDenomIndexKey struct {
	vs []interface{}
}

func (x BasketBalanceBasketIdBatchDenomIndexKey) id() uint32             { return 0 }
func (x BasketBalanceBasketIdBatchDenomIndexKey) values() []interface{}  { return x.vs }
func (x BasketBalanceBasketIdBatchDenomIndexKey) basketBalanceIndexKey() {}

func (this BasketBalanceBasketIdBatchDenomIndexKey) WithBasketId(basket_id uint64) BasketBalanceBasketIdBatchDenomIndexKey {
	this.vs = []interface{}{basket_id}
	return this
}

func (this BasketBalanceBasketIdBatchDenomIndexKey) WithBasketIdBatchDenom(basket_id uint64, batch_denom string) BasketBalanceBasketIdBatchDenomIndexKey {
	this.vs = []interface{}{basket_id, batch_denom}
	return this
}

type BasketBalanceBasketIdBatchStartDateIndexKey struct {
	vs []interface{}
}

func (x BasketBalanceBasketIdBatchStartDateIndexKey) id() uint32             { return 1 }
func (x BasketBalanceBasketIdBatchStartDateIndexKey) values() []interface{}  { return x.vs }
func (x BasketBalanceBasketIdBatchStartDateIndexKey) basketBalanceIndexKey() {}

func (this BasketBalanceBasketIdBatchStartDateIndexKey) WithBasketId(basket_id uint64) BasketBalanceBasketIdBatchStartDateIndexKey {
	this.vs = []interface{}{basket_id}
	return this
}

func (this BasketBalanceBasketIdBatchStartDateIndexKey) WithBasketIdBatchStartDate(basket_id uint64, batch_start_date *timestamppb.Timestamp) BasketBalanceBasketIdBatchStartDateIndexKey {
	this.vs = []interface{}{basket_id, batch_start_date}
	return this
}

type basketBalanceStore struct {
	table ormtable.Table
}

func (this basketBalanceStore) Insert(ctx context.Context, basketBalance *BasketBalance) error {
	return this.table.Insert(ctx, basketBalance)
}

func (this basketBalanceStore) Update(ctx context.Context, basketBalance *BasketBalance) error {
	return this.table.Update(ctx, basketBalance)
}

func (this basketBalanceStore) Save(ctx context.Context, basketBalance *BasketBalance) error {
	return this.table.Save(ctx, basketBalance)
}

func (this basketBalanceStore) Delete(ctx context.Context, basketBalance *BasketBalance) error {
	return this.table.Delete(ctx, basketBalance)
}

func (this basketBalanceStore) Has(ctx context.Context, basket_id uint64, batch_denom string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, basket_id, batch_denom)
}

func (this basketBalanceStore) Get(ctx context.Context, basket_id uint64, batch_denom string) (*BasketBalance, error) {
	var basketBalance BasketBalance
	found, err := this.table.PrimaryKey().Get(ctx, &basketBalance, basket_id, batch_denom)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &basketBalance, nil
}

func (this basketBalanceStore) List(ctx context.Context, prefixKey BasketBalanceIndexKey, opts ...ormlist.Option) (BasketBalanceIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BasketBalanceIterator{it}, err
}

func (this basketBalanceStore) ListRange(ctx context.Context, from, to BasketBalanceIndexKey, opts ...ormlist.Option) (BasketBalanceIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BasketBalanceIterator{it}, err
}

func (this basketBalanceStore) DeleteBy(ctx context.Context, prefixKey BasketBalanceIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this basketBalanceStore) DeleteRange(ctx context.Context, from, to BasketBalanceIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this basketBalanceStore) doNotImplement() {}

var _ BasketBalanceStore = basketBalanceStore{}

func NewBasketBalanceStore(db ormtable.Schema) (BasketBalanceStore, error) {
	table := db.GetTable(&BasketBalance{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&BasketBalance{}).ProtoReflect().Descriptor().FullName()))
	}
	return basketBalanceStore{table}, nil
}

type StateStore interface {
	BasketStore() BasketStore
	BasketClassStore() BasketClassStore
	BasketBalanceStore() BasketBalanceStore

	doNotImplement()
}

type stateStore struct {
	basket        BasketStore
	basketClass   BasketClassStore
	basketBalance BasketBalanceStore
}

func (x stateStore) BasketStore() BasketStore {
	return x.basket
}

func (x stateStore) BasketClassStore() BasketClassStore {
	return x.basketClass
}

func (x stateStore) BasketBalanceStore() BasketBalanceStore {
	return x.basketBalance
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	basketStore, err := NewBasketStore(db)
	if err != nil {
		return nil, err
	}

	basketClassStore, err := NewBasketClassStore(db)
	if err != nil {
		return nil, err
	}

	basketBalanceStore, err := NewBasketBalanceStore(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		basketStore,
		basketClassStore,
		basketBalanceStore,
	}, nil
}
