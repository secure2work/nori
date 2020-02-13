package graph_test

import (
	"fmt"
	"testing"

	"github.com/nori-io/nori-common/v2/meta"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/nori/internal/dependency"
	"github.com/nori-io/nori/pkg/errors"
)

const (
	pluginOne     = "plugin1"
	pluginTwo     = "plugin2"
	pluginThree   = "plugin3"
	pluginFour    = "plugin4"
	pluginRingOne = "pluginRingOne"
	pluginRingTwo = "pluginRingTwo"
	pluginHttp    = "pluginHttp"
	pluginSql     = "pluginSql"
	pluginCms     = "pluginCms"

	InterfaceOne     = meta.Interface("nori/InterfaceOne@1.0.0")
	InterfaceTwo     = meta.Interface("nori/InterfaceTwo@1.0.0")
	InterfaceThree   = meta.Interface("nori/InterfaceThree@1.0.0")
	InterfaceFour    = meta.Interface("nori/InterfaceFour@1.0.0")
	InterfaceRingOne = meta.Interface("nori/RingOne@1.0.0")
	InterfaceRingTwo = meta.Interface("nori/RingTwo@1.0.0")
	InterfaceCms     = meta.Interface("nori/InterfaceCms@1.0.0")
	InterfaceHttp    = meta.Interface("nori/InterfaceHttp@1.0.0")
	InterfaceSql     = meta.Interface("nori/InterfaceSql@1.0.0")
)

// Interface with SelfRing
func plugin_RingOne(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginRingOne,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingOne}},
		Interface:    InterfaceRingOne,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

// Interface with SelfRing
func plugin_RingTwo(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginRingTwo,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingTwo}},
		Interface:    InterfaceRingTwo,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

// plugin1 depends on InterfaceHttp
func plugin1(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginOne,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceHttp}},
		Interface:    InterfaceOne,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

//depend of InterfaceThree
func plugin2(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginTwo,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Dependencies: []meta.Dependency{
			{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree},
		},
		Interface: InterfaceTwo,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

// without dependencies
func plugin3(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginThree,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: InterfaceThree,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

//without dependencies
func plugin4(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginFour,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: InterfaceFour,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

//without dependencies
func plugin_HTTP(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginHttp,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: InterfaceHttp,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

// without dependencies
func plugin_Sql(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginSql,
			Version: "1.0.0",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: InterfaceSql,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

// depend of  interfaceHttp, interfaceSql
func plugin_Cms(deps ...meta.Dependency) meta.Meta {
	data := meta.Data{
		ID: meta.ID{
			ID:      pluginCms,
			Version: "1.0.0",
		},
		Dependencies: []meta.Dependency{
			{
				Constraint: ">=1.0.0, <2.0.0",
				Interface:  InterfaceHttp,
			},
			{
				Constraint: ">=1.0.0, <2.0.0",
				Interface:  InterfaceSql,
			},
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0.0, <2.0.0",
		},
		Interface: InterfaceCms,
	}
	if len(deps) > 0 {
		data.Dependencies = deps
	}
	return data
}

//1) InterfaceOne, InterfaceTwo, InterfaceThree, interfaceHTTP (all available) order for adding - plugin1, plugin3, plugin2
func TestDependencyGraph_AllInterfacesAvailable(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1()))
	a.Nil(managerPlugin.Add(plugin3()))
	a.Nil(managerPlugin.Add(plugin2()))
	a.Nil(managerPlugin.Add(plugin_HTTP()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	pluginsSorted, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}
	t.Log("Plugins' order after sorting:")
	for index, _ := range pluginsSorted {
		t.Log("Plugin n.", index+1, " in list for start:", pluginsSorted[index].String())
	}
	var (
		index1    int
		index2    int
		index3    int
		indexHttp int
	)
	for index, value := range pluginsSorted {
		if value.ID == pluginOne {
			index1 = index
		}
		if value.ID == pluginTwo {
			index2 = index
		}
		if value.ID == pluginThree {
			index3 = index
		}
		if value.ID == pluginHttp {
			indexHttp = index
		}
	}
	fmt.Println(indexHttp)
	a.Equal(true, (indexHttp == 0) || (indexHttp == 1) || (indexHttp == 2))
	a.Equal(true, (index3 == 0) || (index3 == 1) || (index3 == 2))
	a.Equal(true, (index2 == 1) || (index2 == 2) || (index2 == 3))
	a.Equal(true, (index1 == 1) || (index1 == 2) || (index1 == 3))
	a.Equal(true, indexHttp < index1)
	a.Equal(true, index3 < index2)
	a.Equal(4, len(pluginsSorted))
}

//2) plugin1, plugin2, plugin3, pluginHttp (plugins with InterfaceThree and InterfaceHttp are unavailable)
func TestDependencyGraph_UnavailablePlugin3PluginHttp(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1()))
	a.Nil(managerPlugin.Add(plugin2()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)
}

//3) plugin1, plugin3 (plugin that implements interfacesHttp is unavailable)
func TestDependencyGraph_UnavailablePlugin2InterfacesHttp(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1()))
	a.Nil(managerPlugin.Add(plugin3()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)
}

//4) plugin1 -> interfaceHttp (all available)
func TestDependencyGraph_AllPluginsAvailableWithInterfaceDependency(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceHttp,
	})))
	a.Nil(managerPlugin.Add(plugin_HTTP()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sorting:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	pluginsSorted, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}
	t.Log("Plugins' order after sorting:")
	for index, _ := range pluginsSorted {
		t.Log("Plugin n.", index+1, " in list for start:", pluginsSorted[index].String())
	}

	var (
		index1    int
		indexHTTP int
	)
	for index, value := range pluginsSorted {
		if value.ID == pluginOne {
			index1 = index
		}
		if value.ID == "http" {
			indexHTTP = index
		}
	}

	a.Equal(true, indexHTTP < index1)
	a.Equal(2, len(pluginsSorted))
}

//5) plugin1-> interfaceHttp (interface is unavailable)
func TestDependencyGraph_UnavailableInterface(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceHttp,
	})))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)

}

//6) plugin 3 -> interfaceTwo, plugin 2 -> InterfaceFour (all available)
func TestDependencyGraph_AllPluginsAvailable2(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1()))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceFour})))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo})))
	a.Nil(managerPlugin.Add(plugin4()))
	a.Nil(managerPlugin.Add(plugin_HTTP()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	pluginsSorted, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}
	t.Log("Plugins' order after sorting:")
	for index, _ := range pluginsSorted {
		t.Log("Plugin n.", index+1, " in list for start:", pluginsSorted[index].String())
	}
	var (
		index1    int
		index2    int
		index3    int
		index4    int
		indexHttp int
	)
	for index, value := range pluginsSorted {
		if value.ID == pluginOne {
			index1 = index
		}
		if value.ID == pluginTwo {
			index2 = index
		}
		if value.ID == pluginThree {
			index3 = index
		}
		if value.ID == pluginFour {
			index4 = index
		}
		if value.ID == pluginHttp {
			indexHttp = index
		}
	}
	a.Equal(err, nil)
	a.Equal(true, (index4 == 0) || (index4 == 1) || (index4 == 2) || (index4 == 3))
	a.Equal(true, index2 < index3)
	a.Equal(true, indexHttp < index1)
	a.Equal(true, index4 < index2)
	a.Equal(5, len(pluginsSorted))
}

//7) plugin1 -> plugin2, plugin 3 -> plugin2, (plugins with InterfaceHttp and interfaceTwo are unavailable)
func TestDependencyGraph_UnavailablePlugin(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1()))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo})))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)

}

//8) pluginCms->InterfaceSql, pluginCms->interfaceHttp
func TestDependencyGraph_PluginsCmsMySqlHttp(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin_Cms()))
	a.Nil(managerPlugin.Add(plugin_HTTP()))
	a.Nil(managerPlugin.Add(plugin_Sql()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	pluginsSorted, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}
	t.Log("Plugins' order after sorting:")
	for index, _ := range pluginsSorted {
		t.Log("Plugin n.", index+1, " in list for start:", pluginsSorted[index].String())
	}
	var (
		index1Http  int
		index2Mysql int
		index3Cms   int
	)
	for index, value := range pluginsSorted {
		if value.ID == "pluginHTTP" {
			index1Http = index
		}
		if value.ID == "pluginMysql" {
			index2Mysql = index
		}
		if value.ID == "pluginCms" {
			index3Cms = index
		}
	}
	a.Equal(true, index3Cms > index1Http)
	a.Equal(true, index3Cms > index2Mysql)
	a.Equal(3, len(pluginsSorted))
}

//9) ring -plugin1->interfaceOne, plugin2->InterfaceTwo
func TestDependencyGraph_Ring1(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Equal(errors.LoopVertexFound{Dependency: struct {
		Constraint string
		Interface  meta.Interface
	}{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingOne}}, managerPlugin.Add(plugin_RingOne(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingOne})))
	a.Equal(errors.LoopVertexFound{Dependency: struct {
		Constraint string
		Interface  meta.Interface
	}{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingTwo}}, managerPlugin.Add(plugin_RingTwo(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceRingTwo})))

}

//10)ring plugin2->interfaceThree, plugin3->interfaceTwo
func TestDependencyGraph_Ring2(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin2()))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo})))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)

}

//11)ring plugin1->interfaceTwo, plugin2->interfaceThree, plugin3->interfaceTwo, plugin3->interfaceOne
func TestDependencyGraph_Ring3(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceTwo,
	})))
	a.Nil(managerPlugin.Add(plugin2()))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo},
		meta.Dependency{
			Constraint: ">=1.0.0, <2.0.0",
			Interface:  InterfaceOne,
		})))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)

}

//12) plugin1 ->InterfaceTwo, InterfaceHttp,
//pluginHttp-> InterfaceThree, InterfaceSql
//plugin2 -> InterfaceThree
//pluginSql->InterfaceThree
//pluginCms->InterfaceHttp, InterfaceSql
func TestDependencyGraph_Sort1(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo},
		meta.Dependency{
			Constraint: ">=1.0.0, <2.0.0",
			Interface:  InterfaceHttp,
		})))
	a.Nil(managerPlugin.Add(plugin_HTTP(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree},
		meta.Dependency{
			Constraint: ">=1.0.0, <2.0.0",
			Interface:  InterfaceSql,
		})))
	a.Nil(managerPlugin.Add(plugin3()))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Sql(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Cms()))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	pluginsSorted, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}
	t.Log("Plugins' order after sorting:")
	for index, _ := range pluginsSorted {
		t.Log("Plugin n.", index+1, " in list for start:", pluginsSorted[index].String())
	}
	var (
		index1    int
		index2    int
		index3    int
		indexHTTP int
		indexSql  int
		indexCms  int
	)
	for index, value := range pluginsSorted {
		if value.ID == pluginOne {
			index1 = index
		}
		if value.ID == pluginTwo {
			index2 = index
		}
		if value.ID == pluginThree {
			index3 = index
		}
		if value.ID == pluginCms {
			indexCms = index
		}
		if value.ID == pluginHttp {
			indexHTTP = index
		}
		if value.ID == pluginSql {
			indexSql = index
		}
	}
	a.Equal(true, index3 < indexSql)
	a.Equal(true, indexHTTP < indexCms)
	a.Equal(true, indexSql < indexCms)
	a.Equal(true, index2 < index1)
	a.Equal(true, indexHTTP < index1)
	a.Equal(true, index3 < indexHTTP)
	a.Equal(true, indexSql < indexHTTP)
	a.Equal(true, index3 < index2)
	a.Equal(6, len(pluginsSorted))
}

//13) plugin1 ->InterfaceTwo, InterfaceHttp,
//pluginHttp-> InterfaceThree, InterfaceSql
//plugin2 -> InterfaceThree
//pluginSql->InterfaceThree
//pluginCms->InterfaceCms
func TestDependencyGraph_SortWithRing(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo},
		meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceHttp})))
	a.Nil(managerPlugin.Add(plugin_HTTP(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin3()))
	a.Nil(managerPlugin.Add(plugin2()))
	a.Nil(managerPlugin.Add(plugin_Sql(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Equal(errors.LoopVertexFound{Dependency: struct {
		Constraint string
		Interface  meta.Interface
	}{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceCms}}, managerPlugin.Add(plugin_Cms(meta.Dependency{">=1.0.0, <2.0.0", InterfaceCms})))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	_, err := managerPlugin.Sort()
	a.NoError(err)
	t.Log(err)

}

//14) plugin1 ->InterfaceTwo, InterfaceHttp,
//pluginHttp-> InterfaceThree, InterfaceSql
//plugin3->InterfaceHttp
//plugin2 -> InterfaceThree
//pluginSql->InterfaceThree
//pluginCms->InterfaceHttp, InterfaceSql
func TestDependencyGraph_SortWithRing2(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo},
		meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceHttp})))
	a.Nil(managerPlugin.Add(plugin_HTTP(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceHttp})))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Sql(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Cms()))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)
}

//15) plugin1 ->InterfaceTwo, InterfaceHttp,
//pluginHttp-> InterfaceThree
//plugin2 -> InterfaceThree
//plugin3-> InterfaceCms
//pluginSql->InterfaceThree
//pluginCms->InterfaceHttp, InterfaceSql
//pluginHttp->InterfaceThree
func TestDependencyGraph_SortWithRing3(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceTwo},
		meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceHttp})))
	a.Nil(managerPlugin.Add(plugin_HTTP(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceCms})))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Sql(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceThree})))
	a.Nil(managerPlugin.Add(plugin_Cms()))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
	t.Log(err)
}

// ring through InterfaceOne and InterfaceThree
//16) plugin1 -> InterfaceTwo, plugin2 -> InterfaceThree, plugin3->InterfaceOne (all available) order for adding - 1 2 3
func TestDependencyGraph_SortWithRing4(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceTwo,
	})))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceThree,
	})))
	a.Nil(managerPlugin.Add(plugin3(meta.Dependency{Constraint: ">=1.0.0, <2.0.0", Interface: InterfaceOne})))

	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	a.Error(err, "Error in sorting")
}

//17)plugin1-> InterfaceTwo, plugin2->InterfaceOne, plugin3 (all available) order for adding - plugin1, plugin3, plugin2
func TestDependencyGraph_SortWithRing5(t *testing.T) {
	a := assert.New(t)
	managerPlugin := dependency.NewManager()
	a.Nil(managerPlugin.Add(plugin1(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceTwo,
	})))
	a.Nil(managerPlugin.Add(plugin2(meta.Dependency{
		Constraint: ">=1.0.0, <2.0.0",
		Interface:  InterfaceOne,
	})))
	a.Nil(managerPlugin.Add(plugin3()))
	t.Log("Plugins' order until sorting:")
	pluginsList := managerPlugin.GetPluginsList()
	i := 0
	for _, value := range pluginsList {
		i++
		if len(value.GetDependencies()) > 0 {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), " Dependencies:")
			j := 0
			for _, depvalue := range value.GetDependencies() {
				j++
				t.Log("Dependence n.", j, "for", value.Id().ID, "is", depvalue.String())
			}
		} else {
			t.Log("Plugin n.", i, " in list until sotring:", value.Id(), "Plugin doesn't have dependencies")
		}
	}
	t.Log()
	_, err := managerPlugin.Sort()
	if err != nil {
		t.Log("Error in sorting:", err.Error())
	}

}
