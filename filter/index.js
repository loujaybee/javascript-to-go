
const people = [{
    name: "Grace",
    age: 12
}, {
    name: "Bob",
    age: 23
}, {
    name: "Alice",
    age: 46
}]

console.log(
    people.filter((item) => { return item.age > 18 })
)
