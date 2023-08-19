package iter_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
	"github.com/BooleanCat/go-functional/iter/filters"
)

func ExampleLift() {
	positives := iter.Filter[int](iter.Lift([]int{-1, 4, 6, 4, -5}), filters.GreaterThan(-1))
	fmt.Println(positives.Collect())
	// Output: [4 6 4]
}

func TestLift(t *testing.T) {
	items := iter.Lift([]int{1, 2})
	assert.Equal(t, items.Next().Unwrap(), 1)
	assert.Equal(t, items.Next().Unwrap(), 2)
	assert.True(t, items.Next().IsNone())
}

func TestLiftEmpty(t *testing.T) {
	assert.True(t, iter.Lift([]int{}).Next().IsNone())
}

func TestLiftCollect(t *testing.T) {
	items := iter.Lift([]int{1, 2}).Collect()
	assert.SliceEqual(t, items, []int{1, 2})
}

func TestLiftDrop(t *testing.T) {
	items := iter.Lift([]int{1, 2, 3}).Drop(1).Collect()
	assert.SliceEqual(t, items, []int{2, 3})
}

func TestLiftTake(t *testing.T) {
	items := iter.Lift([]int{1, 2, 3}).Take(2).Collect()
	assert.SliceEqual(t, items, []int{1, 2})
}

func TestLiftHashMap(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.Collect[iter.Tuple[string, string]](iter.LiftHashMap(pokemon))
	sort.Slice(items, func(i, j int) bool {
		return items[i].One < items[j].One
	})

	assert.SliceEqual(t, items, []iter.Tuple[string, string]{{"name", "pikachu"}, {"type", "electric"}})
}

func TestLiftHashMapCloseEarly(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMap(pokemon)
	assert.True(t, items.Next().IsSome())
	items.Close()
	assert.True(t, items.Next().IsNone())
}

func TestLiftHashMapCloseMultipleSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMap(pokemon)
	items.Close()
	items.Close()
}

func TestLiftHashMapCloseAfterExhaustedSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMap(pokemon)
	iter.Collect[iter.Tuple[string, string]](items)
	items.Close()
}

func TestLiftHashMapCollect(t *testing.T) {
	items := iter.LiftHashMap(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Collect()

	sort.Slice(items, func(i, j int) bool {
		return items[i].One < items[j].One
	})

	assert.SliceEqual(t, items, []iter.Tuple[string, string]{{"name", "pikachu"}, {"type", "electric"}})
}

func TestLiftHashMapDrop(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMap(pokemon).Drop(1).Collect()

	assert.Equal(t, 1, len(items))
}

func TestLiftHashMapTake(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMap(pokemon).Take(1).Collect()

	assert.Equal(t, 1, len(items))
}

func TestLiftHashMapKeys(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	keys := iter.Collect[string](iter.LiftHashMapKeys(pokemon))
	sort.Strings(keys)

	assert.SliceEqual(t, keys, []string{"name", "type"})
}

func TestLiftHashMapKeysExhausted(t *testing.T) {
	pokemon := iter.LiftHashMapKeys(make(map[string]string))

	iter.Collect[string](pokemon)
	assert.True(t, pokemon.Next().IsNone())
}

func TestLiftHashMapKeysCloseEarly(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapKeys(pokemon)
	assert.True(t, items.Next().IsSome())
	items.Close()
	assert.True(t, items.Next().IsNone())
}

func TestLiftHashMapKeysCloseMultipleSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapKeys(pokemon)
	items.Close()
	items.Close()
}

func TestLiftHashMapKeysCloseAfterExhaustedSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapKeys(pokemon)
	iter.Collect[string](items)
	items.Close()
}

func TestLiftHashMapKeysCollect(t *testing.T) {
	keys := iter.LiftHashMapKeys(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Collect()

	sort.Strings(keys)

	assert.SliceEqual(t, keys, []string{"name", "type"})
}

func TestLiftHashMapKeysDrop(t *testing.T) {
	keys := iter.LiftHashMapKeys(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Drop(1).Collect()

	assert.Equal(t, 1, len(keys))
}

func TestLiftHashMapKeysTake(t *testing.T) {
	keys := iter.LiftHashMapKeys(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Take(1).Collect()

	assert.Equal(t, 1, len(keys))
}

func TestLiftHashMapValues(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	keys := iter.Collect[string](iter.LiftHashMapValues(pokemon))
	sort.Strings(keys)

	assert.SliceEqual(t, keys, []string{"electric", "pikachu"})
}

func TestLiftHashMapValuesExhausted(t *testing.T) {
	pokemon := iter.LiftHashMapValues(make(map[string]string))

	iter.Collect[string](pokemon)
	assert.True(t, pokemon.Next().IsNone())
}

func TestLiftHashMapValuesCloseEarly(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapValues(pokemon)
	assert.True(t, items.Next().IsSome())
	items.Close()
	assert.True(t, items.Next().IsNone())
}

func TestLiftHashMapValuesCloseMultipleSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapValues(pokemon)
	items.Close()
	items.Close()
}

func TestLiftHashMapValuesCloseAfterExhaustedSafe(t *testing.T) {
	pokemon := make(map[string]string)
	pokemon["name"] = "pikachu"
	pokemon["type"] = "electric"

	items := iter.LiftHashMapValues(pokemon)
	iter.Collect[string](items)
	items.Close()
}

func TestLiftHashMapValuesCollect(t *testing.T) {
	values := iter.LiftHashMapValues(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Collect()

	sort.Strings(values)

	assert.SliceEqual(t, values, []string{"electric", "pikachu"})
}

func TestLiftHashMapValuesDrop(t *testing.T) {
	values := iter.LiftHashMapValues(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Drop(1).Collect()

	assert.Equal(t, 1, len(values))
}

func TestLiftHashMapValuesTake(t *testing.T) {
	values := iter.LiftHashMapValues(map[string]string{
		"name": "pikachu",
		"type": "electric",
	}).Take(1).Collect()

	assert.Equal(t, 1, len(values))
}
