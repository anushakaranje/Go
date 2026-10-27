// package main

// import "fmt"

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")
// 	fmt.Println("Starting Pod Health Checker...")
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	podName := "frontend-pod"
// 	podStatus := "Running"

// 	fmt.Println("Pod Name:", podName)
// 	fmt.Println("Pod Status:", podStatus)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []string{
// 		"frontend-pod",
// 		"backend-pod",
// 		"database-pod",
// 	}

// 	fmt.Println("Pods in the application:")

// 	for _, pod := range pods {
// 		fmt.Println(pod)
// 	}
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// }

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running"},
// 		{Name: "backend-pod", Status: "Running"},
// 		{Name: "database-pod", Status: "Pending"},
// 		{Name: "payment-pod", Status: "Failed"},
// 	}

// 	fmt.Printf("%-20s %s\n", "Pod Name", "Status")
// 	fmt.Println("-----------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-20s %s\n", pod.Name, pod.Status)
// 	}
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// }

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running"},
// 		{Name: "backend-pod", Status: "Running"},
// 		{Name: "database-pod", Status: "Pending"},
// 		{Name: "payment-pod", Status: "Failed"},
// 	}

// 	fmt.Printf("%-20s %s\n", "Pod Name", "Status")
// 	fmt.Println("--------------------------------")

// 	healthyPods := 0
// 	unhealthyPods := 0

// 	for _, pod := range pods {
// 		fmt.Printf("%-20s %s\n", pod.Name, pod.Status)

// 		if pod.Status == "Running" {
// 			healthyPods++
// 		} else {
// 			unhealthyPods++
// 		}
// 	}

// 	fmt.Println("--------------------------------")
// 	fmt.Printf("%-20s %d\n", "Total Pods:", len(pods))
// 	fmt.Printf("%-20s %d\n", "Healthy Pods:", healthyPods)
// 	fmt.Printf("%-20s %d\n", "Unhealthy Pods:", unhealthyPods)
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// }

// func checkPodHealth(pods []Pod) (int, int) {
// 	healthyPods := 0
// 	unhealthyPods := 0

// 	for _, pod := range pods {
// 		if pod.Status == "Running" {
// 			healthyPods++
// 		} else {
// 			unhealthyPods++
// 		}
// 	}

// 	return healthyPods, unhealthyPods
// }

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running"},
// 		{Name: "backend-pod", Status: "Running"},
// 		{Name: "database-pod", Status: "Pending"},
// 		{Name: "payment-pod", Status: "Failed"},
// 	}

// 	fmt.Printf("%-20s %s\n", "Pod Name", "Status")
// 	fmt.Println("--------------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-20s %s\n", pod.Name, pod.Status)
// 	}

// 	healthyPods, unhealthyPods := checkPodHealth(pods)

// 	fmt.Println("--------------------------------")
// 	fmt.Printf("%-20s %d\n", "Total Pods:", len(pods))
// 	fmt.Printf("%-20s %d\n", "Healthy Pods:", healthyPods)
// 	fmt.Printf("%-20s %d\n", "Unhealthy Pods:", unhealthyPods)
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pods []Pod) (int, int) {
// 	healthyPods := 0
// 	unhealthyPods := 0

// 	for _, pod := range pods {
// 		if pod.Status == "Running" && pod.Ready {
// 			healthyPods++
// 		} else {
// 			unhealthyPods++
// 		}
// 	}

// 	return healthyPods, unhealthyPods
// }

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running", Ready: true},
// 		{Name: "backend-pod", Status: "Running", Ready: true},
// 		{Name: "database-pod", Status: "Running", Ready: false},
// 		{Name: "payment-pod", Status: "Failed", Ready: false},
// 	}

// 	fmt.Printf("%-20s %-12s %s\n", "Pod Name", "Status", "Ready")
// 	fmt.Println("---------------------------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-20s %-12s %t\n", pod.Name, pod.Status, pod.Ready)
// 	}

// 	healthyPods, unhealthyPods := checkPodHealth(pods)

// 	fmt.Println("---------------------------------------------")
// 	fmt.Printf("%-20s %d\n", "Total Pods:", len(pods))
// 	fmt.Printf("%-20s %d\n", "Healthy Pods:", healthyPods)
// 	fmt.Printf("%-20s %d\n", "Unhealthy Pods:", unhealthyPods)
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pod Pod) string {
// 	if pod.Status == "Failed" {
// 		return "Pod has failed"
// 	}

// 	if pod.Status == "Pending" {
// 		return "Pod is pending"
// 	}

// 	if pod.Status != "Running" {
// 		return "Unknown pod status"
// 	}

// 	if !pod.Ready {
// 		return "Pod is not ready"
// 	}

// 	return "Healthy"
// }

// func main() {
// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")

// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running", Ready: true},
// 		{Name: "backend-pod", Status: "Running", Ready: true},
// 		{Name: "database-pod", Status: "Pending", Ready: false},
// 		{Name: "payment-pod", Status: "Failed", Ready: false},
// 		{Name: "cache-pod", Status: "Running", Ready: false},
// 	}

// 	fmt.Printf("%-18s %-10s %-7s %s\n", "Pod Name", "Status", "Ready", "Health")
// 	fmt.Println("----------------------------------------------------------")

// 	healthyPods := 0
// 	unhealthyPods := 0

// 	for _, pod := range pods {
// 		health := checkPodHealth(pod)

// 		fmt.Printf("%-18s %-10s %-7t %s\n",
// 			pod.Name, pod.Status, pod.Ready, health)

// 		if health == "Healthy" {
// 			healthyPods++
// 		} else {
// 			unhealthyPods++
// 		}
// 	}

// 	fmt.Println("----------------------------------------------------------")
// 	fmt.Printf("%-20s %d\n", "Total Pods:", len(pods))
// 	fmt.Printf("%-20s %d\n", "Healthy Pods:", healthyPods)
// 	fmt.Printf("%-20s %d\n", "Unhealthy Pods:", unhealthyPods)
// }

// package main

// import "fmt"

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pod Pod) string {
// 	if pod.Status == "Failed" {
// 		return "Pod has failed"
// 	}

// 	if pod.Status == "Pending" {
// 		return "Pod is pending"
// 	}

// 	if pod.Status != "Running" {
// 		return "Unknown pod status"
// 	}

// 	if !pod.Ready {
// 		return "Pod is not ready"
// 	}

// 	return "Healthy"
// }

// func showUnhealthyPods(pods []Pod) {
// 	fmt.Println("\nUNHEALTHY PODS")
// 	fmt.Println("---------------------------------------------")

// 	found := false

// 	for _, pod := range pods {
// 		health := checkPodHealth(pod)

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pod Pod) string {
// 	if pod.Status == "Failed" {
// 		return "Pod has failed"
// 	}

// 	if pod.Status == "Pending" {
// 		return "Pod is pending"
// 	}

// 	if pod.Status != "Running" {
// 		return "Unknown pod status"
// 	}

// 	if !pod.Ready {
// 		return "Pod is not ready"
// 	}

// 	return "Healthy"
// }

// func showAllPods(pods []Pod) {
// 	fmt.Printf("%-18s %-10s %-7s %s\n",
// 		"Pod Name", "Status", "Ready", "Health")
// 	fmt.Println("----------------------------------------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-18s %-10s %-7t %s\n",
// 			pod.Name, pod.Status, pod.Ready, checkPodHealth(pod))
// 	}
// }

// func showUnhealthyPods(pods []Pod) {
// 	fmt.Println("\nUNHEALTHY PODS")
// 	fmt.Println("---------------------------------------------")

// 	found := false

// 	for _, pod := range pods {
// 		health := checkPodHealth(pod)

// 		if health != "Healthy" {
// 			fmt.Printf("%-18s %-10s %s\n",
// 				pod.Name, pod.Status, health)
// 			found = true
// 		}
// 	}

// 	if !found {
// 		fmt.Println("All pods are healthy!")
// 	}
// }

// func main() {
// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running", Ready: true},
// 		{Name: "backend-pod", Status: "Running", Ready: true},
// 		{Name: "database-pod", Status: "Pending", Ready: false},
// 		{Name: "payment-pod", Status: "Failed", Ready: false},
// 		{Name: "cache-pod", Status: "Running", Ready: false},
// 	}

// 	fmt.Println("================================")
// 	fmt.Println("       POD HEALTH CHECKER")
// 	fmt.Println("================================")
// 	fmt.Println("1. Show all pods")
// 	fmt.Println("2. Show unhealthy pods")
// 	fmt.Print("Enter your choice (1 or 2): ")

// 	reader := bufio.NewReader(os.Stdin)
// 	choice, err := reader.ReadString('\n')

// 	if err != nil && len(choice) == 0 {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	choice = strings.TrimSpace(choice)

// 	switch choice {
// 	case "1":
// 		showAllPods(pods)
// 	case "2":
// 		showUnhealthyPods(pods)
// 	default:
// 		fmt.Println("Invalid choice. Please enter 1 or 2.")
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pod Pod) string {
// 	if pod.Status == "Failed" {
// 		return "Pod has failed"
// 	}

// 	if pod.Status == "Pending" {
// 		return "Pod is pending"
// 	}

// 	if pod.Status != "Running" {
// 		return "Unknown pod status"
// 	}

// 	if !pod.Ready {
// 		return "Pod is not ready"
// 	}

// 	return "Healthy"
// }

// func showAllPods(pods []Pod) {
// 	fmt.Printf("%-18s %-10s %-7s %s\n",
// 		"Pod Name", "Status", "Ready", "Health")
// 	fmt.Println("----------------------------------------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-18s %-10s %-7t %s\n",
// 			pod.Name, pod.Status, pod.Ready, checkPodHealth(pod))
// 	}
// }

// func showUnhealthyPods(pods []Pod) {
// 	fmt.Println("\nUNHEALTHY PODS")
// 	fmt.Println("---------------------------------------------")

// 	found := false

// 	for _, pod := range pods {
// 		health := checkPodHealth(pod)

// 		if health != "Healthy" {
// 			fmt.Printf("%-18s %-10s %s\n",
// 				pod.Name, pod.Status, health)
// 			found = true
// 		}
// 	}

// 	if !found {
// 		fmt.Println("All pods are healthy!")
// 	}
// }

// func main() {
// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running", Ready: true},
// 		{Name: "backend-pod", Status: "Running", Ready: true},
// 		{Name: "database-pod", Status: "Pending", Ready: false},
// 		{Name: "payment-pod", Status: "Failed", Ready: false},
// 		{Name: "cache-pod", Status: "Running", Ready: false},
// 	}

// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println("\n================================")
// 		fmt.Println("       POD HEALTH CHECKER")
// 		fmt.Println("================================")
// 		fmt.Println("1. Show all pods")
// 		fmt.Println("2. Show unhealthy pods")
// 		fmt.Println("3. Exit")
// 		fmt.Print("Enter your choice: ")

// 		choice, err := reader.ReadString('\n')

// 		if err != nil && len(choice) == 0 {
// 			fmt.Println("Error reading input:", err)
// 			return
// 		}

// 		choice = strings.TrimSpace(choice)

// 		switch choice {
// 		case "1":
// 			showAllPods(pods)

// 		case "2":
// 			showUnhealthyPods(pods)

// 		case "3":
// 			fmt.Println("Exiting Pod Health Checker...")
// 			return

// 		default:
// 			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
// 		}
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Pod struct {
// 	Name   string
// 	Status string
// 	Ready  bool
// }

// func checkPodHealth(pod Pod) string {
// 	if pod.Status == "Failed" {
// 		return "Pod has failed"
// 	}

// 	if pod.Status == "Pending" {
// 		return "Pod is pending"
// 	}

// 	if pod.Status != "Running" {
// 		return "Unknown pod status"
// 	}

// 	if !pod.Ready {
// 		return "Pod is not ready"
// 	}

// 	return "Healthy"
// }

// func showAllPods(pods []Pod) {
// 	fmt.Printf("%-18s %-10s %-7s %s\n",
// 		"Pod Name", "Status", "Ready", "Health")
// 	fmt.Println("----------------------------------------------------------")

// 	for _, pod := range pods {
// 		fmt.Printf("%-18s %-10s %-7t %s\n",
// 			pod.Name, pod.Status, pod.Ready, checkPodHealth(pod))
// 	}
// }

// func showUnhealthyPods(pods []Pod) {
// 	fmt.Println("\nUNHEALTHY PODS")
// 	fmt.Println("---------------------------------------------")

// 	found := false

// 	for _, pod := range pods {
// 		health := checkPodHealth(pod)

// 		if health != "Healthy" {
// 			fmt.Printf("%-18s %-10s %s\n",
// 				pod.Name, pod.Status, health)
// 			found = true
// 		}
// 	}

// 	if !found {
// 		fmt.Println("All pods are healthy!")
// 	}
// }

// func showHealthSummary(pods []Pod) {
// 	healthy := 0
// 	unhealthy := 0

// 	for _, pod := range pods {
// 		if checkPodHealth(pod) == "Healthy" {
// 			healthy++
// 		} else {
// 			unhealthy++
// 		}
// 	}

// 	fmt.Println("\nHEALTH SUMMARY")
// 	fmt.Println("------------------------")
// 	fmt.Printf("Total Pods:     %d\n", len(pods))
// 	fmt.Printf("Healthy Pods:   %d\n", healthy)
// 	fmt.Printf("Unhealthy Pods: %d\n", unhealthy)
// }

// func main() {
// 	pods := []Pod{
// 		{Name: "frontend-pod", Status: "Running", Ready: true},
// 		{Name: "backend-pod", Status: "Running", Ready: true},
// 		{Name: "database-pod", Status: "Pending", Ready: false},
// 		{Name: "payment-pod", Status: "Failed", Ready: false},
// 		{Name: "cache-pod", Status: "Running", Ready: false},
// 	}

// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Println("\n================================")
// 		fmt.Println("       POD HEALTH CHECKER")
// 		fmt.Println("================================")
// 		fmt.Println("1. Show all pods")
// 		fmt.Println("2. Show unhealthy pods")
// 		fmt.Println("3. Show health summary")
// 		fmt.Println("4. Exit")
// 		fmt.Print("Enter your choice: ")

// 		choice, err := reader.ReadString('\n')

// 		if err != nil && len(choice) == 0 {
// 			fmt.Println("Error reading input:", err)
// 			return
// 		}

// 		choice = strings.TrimSpace(choice)

// 		switch choice {
// 		case "1":
// 			showAllPods(pods)

// 		case "2":
// 			showUnhealthyPods(pods)

// 		case "3":
// 			showHealthSummary(pods)

// 		case "4":
// 			fmt.Println("Exiting Pod Health Checker...")
// 			return

// 		default:
// 			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
// 		}
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pod struct {
	Name   string
	Status string
	Ready  bool
}

func checkPodHealth(pod Pod) string {
	if pod.Status == "Failed" {
		return "Pod has failed"
	}

	if pod.Status == "Pending" {
		return "Pod is pending"
	}

	if pod.Status != "Running" {
		return "Unknown pod status"
	}

	if !pod.Ready {
		return "Pod is not ready"
	}

	return "Healthy"
}

func showAllPods(pods []Pod) {
	fmt.Printf("%-18s %-10s %-7s %s\n",
		"Pod Name", "Status", "Ready", "Health")
	fmt.Println("----------------------------------------------------------")

	for _, pod := range pods {
		fmt.Printf("%-18s %-10s %-7t %s\n",
			pod.Name, pod.Status, pod.Ready, checkPodHealth(pod))
	}
}

func showUnhealthyPods(pods []Pod) {
	fmt.Println("\nUNHEALTHY PODS")
	fmt.Println("---------------------------------------------")

	found := false

	for _, pod := range pods {
		health := checkPodHealth(pod)

		if health != "Healthy" {
			fmt.Printf("%-18s %-10s %s\n",
				pod.Name, pod.Status, health)
			found = true
		}
	}

	if !found {
		fmt.Println("All pods are healthy!")
	}
}

func showHealthSummary(pods []Pod) {
	healthy := 0
	unhealthy := 0

	for _, pod := range pods {
		if checkPodHealth(pod) == "Healthy" {
			healthy++
		} else {
			unhealthy++
		}
	}

	fmt.Println("\nHEALTH SUMMARY")
	fmt.Println("------------------------")
	fmt.Printf("Total Pods:     %d\n", len(pods))
	fmt.Printf("Healthy Pods:   %d\n", healthy)
	fmt.Printf("Unhealthy Pods: %d\n", unhealthy)
}

func main() {
	pods := []Pod{
		{Name: "frontend-pod", Status: "Running", Ready: true},
		{Name: "backend-pod", Status: "Running", Ready: true},
		{Name: "database-pod", Status: "Pending", Ready: false},
		{Name: "payment-pod", Status: "Failed", Ready: false},
		{Name: "cache-pod", Status: "Running", Ready: false},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n================================")
		fmt.Println("       POD HEALTH CHECKER")
		fmt.Println("================================")
		fmt.Println("1. Show all pods")
		fmt.Println("2. Show unhealthy pods")
		fmt.Println("3. Show health summary")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')

		if err != nil && len(choice) == 0 {
			fmt.Println("Error reading input:", err)
			return
		}

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			showAllPods(pods)

		case "2":
			showUnhealthyPods(pods)

		case "3":
			showHealthSummary(pods)

		case "4":
			fmt.Println("Exiting Pod Health Checker...")
			return

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}
}
