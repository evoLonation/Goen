package entity

import (
	"Cocome/entityManager"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInterfaceManager(t *testing.T) {
	item := ItemManager.New()
	item2 := ItemManager.New()
	item.SetPrice(123.0123)
	item2.SetPrice(456.0123)
	item.SetBarcode(114514)
	item2.SetBarcode(1919810)
	item.SetBelongedItem(item2)
	item.AddContainedItem(item2)
	item2.AddContainedItem(item)
	item2.SetBelongedItem(item)

	err := entityManager.Saver.Save()
	require.NoError(t, err)

	item3 := item.GetBelongedItem()
	require.NotNil(t, item3)
	itemArr := item.GetContainedItem()
	require.NotNil(t, itemArr)
}
