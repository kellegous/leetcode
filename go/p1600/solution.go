package p1600

// this question is complete garbage in every imaginable way.

type ThroneInheritance struct {
	name     string
	children map[string][]string
	dead     map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		name:     kingName,
		children: map[string][]string{},
		dead:     map[string]bool{},
	}
}

func (i *ThroneInheritance) Birth(parentName string, childName string) {
	i.children[parentName] = append(i.children[parentName], childName)
}

func (i *ThroneInheritance) Death(name string) {
	i.dead[name] = true

}

func (i *ThroneInheritance) GetInheritanceOrder() []string {
	o := order{
		seen: map[string]bool{},
		inh:  i,
	}
	o.add(i.name)
	return o.name
}

type order struct {
	name []string
	seen map[string]bool
	inh  *ThroneInheritance
}

func (o *order) add(name string) {
	if !o.inh.dead[name] {
		o.name = append(o.name, name)
		o.seen[name] = true
	}

	for _, child := range o.inh.children[name] {
		o.add(child)
	}
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
