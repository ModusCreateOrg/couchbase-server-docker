package main

import (
	"fmt"
	"log"

	"github.com/coreos/go-etcd/etcd"
)

const (
	KEY_CLUSTER_INITIAL_NODE = "cluster-initial-node"
	TTL_NONE                 = 0
)

type CouchbaseCluster struct {
	etcdClient *etcd.Client
}

func (c *CouchbaseCluster) StartCouchbaseNode() error {

	c.etcdClient = etcd.NewClient([]string{"http://127.0.0.1:4001"})
	ok, err := c.BecomeFirstClusterNode()
	if err != nil {
		return err
	}

	fmt.Printf("ok: %v", ok)

	// client.SyncCluster()
	// nodes := client.GetCluster()
	// fmt.Printf("nodes: %+v", nodes)
	return nil
}

func (c CouchbaseCluster) BecomeFirstClusterNode() (bool, error) {

	// TODO: this needs to be passed in as cmd line arg!!
	ip := "127.0.0.1"

	response, err := c.etcdClient.Create(KEY_CLUSTER_INITIAL_NODE, ip, TTL_NONE)

	fmt.Printf("response: %+v, err: %v", response, err)

	if err != nil {
		return false, err
	}

	// check the response

	return true, nil

}

func main() {

	couchbaseCluster := &CouchbaseCluster{}
	if err := couchbaseCluster.StartCouchbaseNode(); err != nil {
		log.Fatal(err)
	}

	/*


		           # Discussion with James:

		           # Use client.write('/cluster-first-node', '<our ip address>', prevExist=False) to
		           # to do an ADD operation.  Only first one to do this will win.

		           # Each node will have a sidekick to maintain its state in
		           # /node-state/<node ip> -> ON or the key will be missing, meaning its not running

		           # When a node comes up:
		           #    - It will try to be the cluster-first-node
		           #    - If succeeds:
		           #       - couchbase-cli cluster-init
		           #       - couchbase-cli bucket-create
		           #    - Else: (there is already a cluster)
		           #       - Find list of machines in cluster from etcd
		           #       - Loop over machines
		           #           - Get ip of machine
		           #           - Does /node-state/<ip> key exist?
		           #              - If it does, try to join the cluster via couchbase-cli server-add





			   # Scenarios

			   # 1. Setting up cluster from scratch (recent machine boot)
			   #
			   #   1A. We become the leader
			   #   1B. Someone beats us to becoming the leader, so we join the leader
			   #
			   # 2. This node rebooted and there is an existing cluster
			   #
			   #   2A. We *were* the leader previously
			   #   2B. We were not the leader previously

			   #  ---- v2

			   # Call etcd to get cluster-leader-counter key

			       # If 0, then we are setting up cluster from scratch (recent machine boot)

			         # success = try_to_become_leader()
			         # if success:
			         #     cluster_one_time_init()
			         #     create_bucket()
			         # else:
			         #     leader = find_leader() && join_leader(leader)

			       # If >= 1, then this node rebooted and there is an existing cluster

			         # leader = find_leader()
			         # if leader == null: abort
			         # join_leader(leader)

			   # try_to_become_leader():

			      # Call etcd to get cluster-leader-ip key

			      # if non-empty, abort and return error

			      # Use previous key index to do a CAS request to write our ip to cluster-leader-ip

			      # if CAS request failed, abort and return error

			      # otherwise, we are the leader

			   # find_leader()

			     # Call etcd to get cluster-leader-ip key or return error

			   # join_leader()

			     #



	*/

	fmt.Sprintf("hello..")

}
