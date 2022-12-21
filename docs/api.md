# API Reference

Packages:

- [v1.edp.epam.com/v1](#v1edpepamcomv1)
- [v1.edp.epam.com/v1alpha1](#v1edpepamcomv1alpha1)

# v1.edp.epam.com/v1

Resource Types:

- [EDPComponent](#edpcomponent)




## EDPComponent
<sup><sup>[↩ Parent](#v1edpepamcomv1 )</sup></sup>






EDPComponent is the Schema for the edpcomponents API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v1.edp.epam.com/v1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>EDPComponent</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#edpcomponentspec">spec</a></b></td>
        <td>object</td>
        <td>
          EDPComponentSpec defines the desired state of EDPComponent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>object</td>
        <td>
          EDPComponentStatus defines the observed state of EDPComponent.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EDPComponent.spec
<sup><sup>[↩ Parent](#edpcomponent)</sup></sup>



EDPComponentSpec defines the desired state of EDPComponent.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>icon</b></td>
        <td>string</td>
        <td>
          base64 encoded SVG icon of a component<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          specifies a type of component, e.g. 'nexus', 'gerrit', etc.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          specifies a link to component<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>visible</b></td>
        <td>boolean</td>
        <td>
          specifies whether a component is visible<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>

# v1.edp.epam.com/v1alpha1

Resource Types:

- [EDPComponent](#edpcomponent)




## EDPComponent
<sup><sup>[↩ Parent](#v1edpepamcomv1alpha1 )</sup></sup>






EDPComponent is the Schema for the edpcomponents API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>v1.edp.epam.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>EDPComponent</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#edpcomponentspec-1">spec</a></b></td>
        <td>object</td>
        <td>
          EDPComponentSpec defines the desired state of EDPComponent.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>status</b></td>
        <td>object</td>
        <td>
          EDPComponentStatus defines the observed state of EDPComponent.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### EDPComponent.spec
<sup><sup>[↩ Parent](#edpcomponent-1)</sup></sup>



EDPComponentSpec defines the desired state of EDPComponent.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>icon</b></td>
        <td>string</td>
        <td>
          base64 encoded SVG icon of a component<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>type</b></td>
        <td>string</td>
        <td>
          specifies a type of component, e.g. 'nexus', 'gerrit', etc.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>url</b></td>
        <td>string</td>
        <td>
          specifies a link to component<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>visible</b></td>
        <td>boolean</td>
        <td>
          specifies whether a component is visible<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>