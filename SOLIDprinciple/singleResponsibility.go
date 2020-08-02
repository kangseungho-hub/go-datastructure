package SOLIDprinciple

// type financeReport struct {
// }

// func (r financeReport) MakeReport() *Report {

// }

// func (r financeReport) SendReport() *Report {
// 	switch{
// 	case email:
// 		...
// 	case json:
// 		...
// 	}
// }
// //+ 이렇게 구현된 경우 확장에는 닫혀있고, 변경에 열려있는 코드가 됨 OCP를 지향하지 않은 구조
// //변경이 열려있다 : 기능을 추가할때 object를 추가하는 것이 아닌 원래 코드를 수정해야 하는 구조

// //두개의 reponsibility를 가지게 되어 S에 위배된다. (프로그램 설계에나 환경에 따라 2개 그 이상을 가질수도 있지만, 가정이다.)
// //Send report 의 전송형식이 Json으로 변경되거나, network패킷으로 변경되는 등 확장과정에서 복잡해진다.
// //따라서 각 Sender객체를 만들어 ReportSender라는 interface 관계를 추가해주고,

// type ReportSender interface {
// 	SendReport(*Report)
// }

// type ReportEmailSender struct{

// }

// func (RES *ReportEmailSender)SendReport(){

// }

// type ReportJsonSender struct{

// }

// func (RJS *ReportJsonSender) SendReport(){

// }

// //이래 Send기능을 분리하여 interface관계를 구현하여 단일 책임을 지향할 수 있음
// //+확장이 열려있고, 변경이 닫혀있어 OCP도 지향
