package main

type TestGun struct {
	Gun
}

func newTestGun() IGun {
	return &TestGun{Gun{
		name:  "TestGun gun",
		power: 5,
	}}
}
