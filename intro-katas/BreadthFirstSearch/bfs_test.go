package bfs

import "testing"

func TestClosestMangoSeller(t *testing.T) {
	cases := []struct {
		contacts       Node
		expectedPerson string
	}{
		{testContacts(), "Thom"},
		{testContactsWithNoSeller(), NotFound},
	}

	for _, c := range cases {
		actualPerson := ClosestMangoSeller(c.contacts)
		if actualPerson != c.expectedPerson {
			t.Errorf("Given %v, expected mango seller of %v, but got %v", c.contacts, c.expectedPerson, actualPerson)
		}
	}
}

func testContacts() Node {
	thom := Node{Value: Person{Name: "Thom", MangoSeller: true}}

	anuj := Node{Value: Person{Name: "Anuj", MangoSeller: false}}

	peggy := Node{Value: Person{Name: "Peggy", MangoSeller: false}}

	bob := Node{
		Value:    Person{Name: "Bob", MangoSeller: false},
		Contacts: []Node{anuj, peggy},
	}

	alice := Node{
		Value:    Person{Name: "Alice", MangoSeller: false},
		Contacts: []Node{peggy},
	}

	jonny := Node{Value: Person{Name: "Jonny", MangoSeller: false}}

	claire := Node{
		Value:    Person{Name: "Claire", MangoSeller: false},
		Contacts: []Node{thom, jonny},
	}

	you := Node{
		Value:    Person{Name: "You", MangoSeller: false},
		Contacts: []Node{bob, claire, alice},
	}

	return you
}

func testContactsWithNoSeller() Node {
	jonny := Node{Value: Person{Name: "Jonny", MangoSeller: false}}

	claire := Node{
		Value:    Person{Name: "Claire", MangoSeller: false},
		Contacts: []Node{jonny},
	}

	you := Node{
		Value:    Person{Name: "You", MangoSeller: false},
		Contacts: []Node{claire},
	}

	return you
}
