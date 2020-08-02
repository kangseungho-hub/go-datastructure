package SOLIDprinciple

// type Actor interface {
// 	Attack()
// 	Move()
// 	Talk()
// }

// type TalkAble interface {
// 	Talk()
// }

// type MoveAble interface {
// 	Move()
// }

// type AttackAble interface {
// 	Attack()
// }

// func MoveTo(a Actor) {
// 	a.Move()
// 	a.Attack()
// }

//이게 interface Segeragation을 위배하는 함수다. 의존성을 가지게되면 이런 함수를 구현하게 될 수도 있기때문에
//Moveable와 같이 interface를 나눠서 아래처럼 위와같은 상황을 강제로 방지할 수 있다.

// func MoveTo(m MoveAble) {
// 	m.Move()
// }
