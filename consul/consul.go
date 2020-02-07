package main

const consulTxnLimit = 64
const maximumPayloadSize = 500000 // max size is actually 512kb

/*
	Reference
	https://github.com/hashicorp/consul/tree/master/api
	https://godoc.org/github.com/hashicorp/consul/api
*/

/*
// Fetch queries the Consul API defined by the given client.
func (d *KVGetQuery) Fetch(clients *ClientSet, opts *QueryOptions) (interface{}, *ResponseMetadata, error) {
	select {
	case <-d.stopCh:
		return nil, nil, ErrStopped
	default:
	}

	opts = opts.Merge(&QueryOptions{
		Datacenter: d.dc,
	})

	log.Printf("[TRACE] %s: GET %s", d, &url.URL{
		Path:     "/v1/kv/" + d.key,
		RawQuery: opts.String(),
	})

	pair, qm, err := clients.Consul().KV().Get(d.key, opts.ToConsulOpts())
	if err != nil {
		return nil, nil, errors.Wrap(err, d.String())
	}

	rm := &ResponseMetadata{
		LastIndex:   qm.LastIndex,
		LastContact: qm.LastContact,
		Block:       d.block,
	}

	if pair == nil {
		log.Printf("[TRACE] %s: returned nil", d)
		return nil, rm, nil
	}

	value := string(pair.Value)
	log.Printf("[TRACE] %s: returned %q", d, value)
	return value, rm, nil
}
*/
