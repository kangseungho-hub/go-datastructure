package SOLIDprinciple

// type Monster struct {
// }

// type Player struct {
// }

// func (p *Player) Attack(m *Monster) {

// }

// //이카면 몬스터가 플레이어를 떄리는 경우, 플레이어가 플레이어를 떄리는 경우 일일이 추가해야하고, 다른 Attackable한 object가 추가될
// //경우 N^2형태로 계속 늘어나게된다.

// type Monster struct {
// }

// type Player struct {
// }

// //Attackable interface관계를 추가해주면 쉽게 해결되는 문제
// type Attackble interface {
// 	Attack()
// }

// type BeAttackable interface {
// 	BeAttacked()
// }

// func (p *Player) Attack(target BeAttackable) {
// 	//...
// 	target.BeAttacked(dmg)
// }

// func Attack(attacker Attackble, target BeAttackable) {
// 	attacker.Attack(target)
// }

//요런 느낌으로 SOLID를 모두 지향하는 것이 깔끔한 코드를 위한 길
