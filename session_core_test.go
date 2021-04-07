package gotdd_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcalbg/gotdd"
	"github.com/alcalbg/gotdd/test/assert"
	"github.com/alcalbg/gotdd/test/doubles"
)

func TestStoreAndRetrieveData(t *testing.T) {

	t.Run("test storing, retrieving, and deleting a simple string", func(t *testing.T) {
		r := &http.Request{}
		w := httptest.NewRecorder()
		sessionStoreSpy := doubles.NewGorillaSessionStoreSpy("")
		ses := gotdd.NewSession(sessionStoreSpy)

		err := ses.Put(w, r, "test", "somevalue")
		assert.NoError(t, err)

		value, err := ses.Get(r, "test")
		assert.NoError(t, err)
		assert.Equal(t, "somevalue", value)

		err = ses.Delete(w, r, "test")
		assert.NoError(t, err)

		value, err = ses.Get(r, "test")
		assert.Error(t, err)
		assert.Equal(t, "", value)
	})

	t.Run("test retrieving unsaved data", func(t *testing.T) {
		r := &http.Request{}
		sessionStoreSpy := doubles.NewGorillaSessionStoreSpy("")
		ses := gotdd.NewSession(sessionStoreSpy)

		value, err := ses.Get(r, "test")
		assert.Error(t, err)
		assert.Equal(t, "", value)
	})

}

func TestDestroyActiveSession(t *testing.T) {

	testUser := doubles.FakeUser1

	r := &http.Request{}
	w := httptest.NewRecorder()
	sessionStoreSpy := doubles.NewGorillaSessionStoreSpy(testUser.GetID())
	ses := gotdd.NewSession(sessionStoreSpy)

	id, err := ses.GetUserID(r)
	assert.NoError(t, err)
	assert.Equal(t, testUser.GetID(), id)

	ses.Destroy(w, r)

	id, err = ses.GetUserID(r)
	assert.Error(t, err)
	assert.Equal(t, "", id)
	assert.Equal(t, 1, sessionStoreSpy.SaveCalls)
	assert.Equal(t, -1, sessionStoreSpy.Session.Options.MaxAge)
}