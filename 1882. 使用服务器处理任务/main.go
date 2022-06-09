package main

import (
	"container/heap"
	"fmt"
)

//给你两个 下标从 0 开始 的整数数组 servers 和 tasks ，长度分别为 n​​​​​​ 和 m​​​​​​ 。servers[i] 是第 i​​​​​​​​​​ 台服务器的 权重 ，而 tasks[j] 是处理第 j​​​​​​ 项任务 所需要的时间（单位：秒）。
//
//你正在运行一个仿真系统，在处理完所有任务后，该系统将会关闭。每台服务器只能同时处理一项任务。第 0 项任务在第 0 秒可以开始处理，相应地，第 j 项任务在第 j 秒可以开始处理。处理第 j 项任务时，你需要为它分配一台 权重最小 的空闲服务器。如果存在多台相同权重的空闲服务器，请选择 下标最小 的服务器。如果一台空闲服务器在第 t 秒分配到第 j 项任务，那么在 t + tasks[j] 时它将恢复空闲状态。
//
//如果没有空闲服务器，则必须等待，直到出现一台空闲服务器，并 尽可能早 地处理剩余任务。 如果有多项任务等待分配，则按照 下标递增 的顺序完成分配。
//
//如果同一时刻存在多台空闲服务器，可以同时将多项任务分别分配给它们。
//
//构建长度为 m 的答案数组 ans ，其中 ans[j] 是第 j 项任务分配的服务器的下标。
//
//返回答案数组 ans​​​​ 。
//
// 
//
//示例 1：
//
//输入：servers = [3,3,2], tasks = [1,2,3,2,1,2]
//输出：[2,2,0,2,1,2]
//解释：事件按时间顺序如下：
//- 0 秒时，第 0 项任务加入到任务队列，使用第 2 台服务器处理到 1 秒。
//- 1 秒时，第 2 台服务器空闲，第 1 项任务加入到任务队列，使用第 2 台服务器处理到 3 秒。
//- 2 秒时，第 2 项任务加入到任务队列，使用第 0 台服务器处理到 5 秒。
//- 3 秒时，第 2 台服务器空闲，第 3 项任务加入到任务队列，使用第 2 台服务器处理到 5 秒。
//- 4 秒时，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 5 秒。
//- 5 秒时，所有服务器都空闲，第 5 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
//示例 2：
//
//输入：servers = [5,1,4,3,2], tasks = [2,1,2,4,5,2,1]
//输出：[1,4,1,4,1,3,2]
//解释：事件按时间顺序如下：
//- 0 秒时，第 0 项任务加入到任务队列，使用第 1 台服务器处理到 2 秒。
//- 1 秒时，第 1 项任务加入到任务队列，使用第 4 台服务器处理到 2 秒。
//- 2 秒时，第 1 台和第 4 台服务器空闲，第 2 项任务加入到任务队列，使用第 1 台服务器处理到 4 秒。
//- 3 秒时，第 3 项任务加入到任务队列，使用第 4 台服务器处理到 7 秒。
//- 4 秒时，第 1 台服务器空闲，第 4 项任务加入到任务队列，使用第 1 台服务器处理到 9 秒。
//- 5 秒时，第 5 项任务加入到任务队列，使用第 3 台服务器处理到 7 秒。
//- 6 秒时，第 6 项任务加入到任务队列，使用第 2 台服务器处理到 7 秒。
// 
//
//提示：
//
//servers.length == n
//tasks.length == m
//1 <= n, m <= 2 * 105
//1 <= servers[i], tasks[j] <= 2 * 105
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/process-tasks-using-servers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {
	//fmt.Println(assignTasks([]int{3,3,2},[]int{1,2,3,2,1,2}))
	fmt.Println(assignTasks([]int{5,1,4,3,2},[]int{2,1,2,4,5,2,1}))
}
type free struct {
	index int
	sort int
}

type freesService struct{
	Frees []*free
}

func (h *freesService) Len() int {
	return len(h.Frees)
}

func (h *freesService) Swap(i, j int) {
	h.Frees[i],h.Frees[j] = h.Frees[j],h.Frees[i]
}

func (h *freesService) Less(i, j int) bool {
	if h.Frees[i].sort ==  h.Frees[j].sort {
		return h.Frees[i].index <  h.Frees[j].index
	}

	return h.Frees[i].sort <  h.Frees[j].sort
}

func (h *freesService) Push(x interface{})   {
	h.Frees = append(h.Frees, x.(*free))
}

func (h *freesService) Pop() (interface{}) {
	res := h.Frees[len(h.Frees) - 1]
	h.Frees = h.Frees[:len(h.Frees) - 1]
	return res
}

type busy struct {
	t int
	index int
	sort int
}

type busysService struct{
	Busys []*busy
}

func (h *busysService) Len() int {
	return len(h.Busys)
}

func (h *busysService) Swap(i, j int) {
	h.Busys[i],h.Busys[j] = h.Busys[j],h.Busys[i]
}

func (h *busysService) Less(i, j int) bool {
	return h.Busys[i].t <  h.Busys[j].t
}

func (h *busysService) Push(x interface{})   {
	h.Busys = append(h.Busys, x.(*busy))
}

func (h *busysService) Pop() (interface{}) {
	res := h.Busys[len(h.Busys) - 1]
	h.Busys = h.Busys[:len(h.Busys) - 1]
	return res
}



var result []int
func assignTasks(servers []int, tasks []int) []int {
	result = make([]int,0,len(tasks))
	freesServices := &freesService{
		Frees: make([]*free,0,len(servers)),
	}
	for k,server := range servers {
		freesServices.Frees = append(freesServices.Frees,&free{
			index: k,
			sort: server,
		})
	}
	heap.Init(freesServices)
	assign(freesServices,&busysService{Busys: make([]*busy,0)},0,0,tasks)

	return result
}

func assign(freesService *freesService,busyService *busysService,t int,index int,tasks []int)  {
	//处理到最后一位
	if index == len(tasks) {
		return
	}

	//先处理busy
	if len(busyService.Busys) >0 && busyService.Busys[0].t <= t{
		busys := heap.Pop(busyService).(*busy)
		heap.Push(freesService,&free{
			index: busys.index,
			sort: busys.sort,
		})
		assign(freesService,busyService,t,index,tasks)
		return
	}

	//没空闲的
	if freesService.Len() ==0 {
		busys := heap.Pop(busyService).(*busy)
		heap.Push(freesService,&free{
			index: busys.index,
			sort: busys.sort,
		})
		assign(freesService,busyService,busys.t,index,tasks)
		return
	}
	//还没到任务开始时间
	if t < index {
		assign(freesService,busyService,index,index,tasks)
		return
	}
	//再处理index

	frees := heap.Pop(freesService).(*free)
	heap.Push(busyService,&busy{
		index: frees.index,
		t:t+tasks[index],
		sort: frees.sort,
	})

	index++
	result = append(result,frees.index)
	assign(freesService,busyService,t,index,tasks)
}
